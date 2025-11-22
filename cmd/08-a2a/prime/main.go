package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	// ADK(Agent Development Kit) 및 관련 라이브러리 임포트
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/web"
	"google.golang.org/adk/cmd/launcher/web/a2a"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/session"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"
	"google.golang.org/genai"
)

// checkPrime은 에이전트가 실제로 호출할 Go 함수입니다.
// tool.Context와 인자 구조체를 받아 소수 여부를 문자열로 반환합니다.
func checkPrime(ctx tool.Context, args struct{ Num int }) (string, error) {
	n := args.Num
	// 1 이하는 소수가 아님
	if n <= 1 {
		return "false", nil
	}
	// 2부터 제곱근까지 나누어 떨어지는지 확인하여 소수 판별
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return "false", nil
		}
	}
	return "true", nil
}

// 추가 함수 1: 팩토리얼 계산
func calculateFactorial(ctx tool.Context, args struct{ N int }) (string, error) {
	n := args.N
	if n < 0 {
		return "", fmt.Errorf("음수는 팩토리얼을 계산할 수 없습니다")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return strconv.Itoa(result), nil
}

// 추가 함수 2: 최대공약수(GCD) 계산 (인자가 2개인 경우)
func calculateGCD(ctx tool.Context, args struct{ A, B int }) (string, error) {
	a, b := args.A, args.B
	for b != 0 {
		a, b = b, a%b
	}
	return strconv.Itoa(a), nil
}

func main() {
	ctx := context.Background()

	// 1. Gemini 모델 초기화
	// 지정된 모델명("gemini-2.5-flash-lite")을 사용하여 클라이언트를 생성합니다.
	model, _ := gemini.NewModel(ctx, "gemini-2.5-flash-lite", &genai.ClientConfig{})

	// 2. 도구(Tool) 생성
	// 기존 소수 판별 도구
	primeTool, _ := functiontool.New(functiontool.Config{
		Name:        "check_prime",
		Description: "Checks if a number is prime",
	}, checkPrime)

	// 팩토리얼 도구 등록
	factorialTool, _ := functiontool.New(functiontool.Config{
		Name:        "calculate_factorial",
		Description: "Calculates the factorial of a number (e.g., 5!)",
	}, calculateFactorial)

	// 최대공약수 도구 등록
	gcdTool, _ := functiontool.New(functiontool.Config{
		Name:        "calculate_gcd",
		Description: "Calculates the Greatest Common Divisor (GCD) of two numbers",
	}, calculateGCD)

	// 3. 에이전트(Agent) 생성 및 도구 목록 업데이트
	mathAgent, _ := llmagent.New(llmagent.Config{
		Name:  "MathHelper", // 이름 변경
		Model: model,
		// 지시문(Instruction)을 업데이트하여 에이전트가 자신의 능력을 알게 합니다.
		Instruction: "You are a helpful math assistant. You can check prime numbers, calculate factorials, and find the GCD of two numbers using the provided tools.",
		// Tools 배열에 새로 만든 도구들을 추가합니다.
		Tools: []tool.Tool{primeTool, factorialTool, gcdTool},
	})

	// 4. 웹 서버 런처 설정
	// A2A(Agent-to-Agent) 통신을 지원하는 웹 런처를 생성합니다.
	port := 8001
	webLauncher := web.NewLauncher(a2a.NewLauncher())

	// 서버 포트 및 A2A 에이전트 URL 설정 (커맨드 라인 인자를 코드로 파싱)
	webLauncher.Parse([]string{
		"--port", strconv.Itoa(port),
		"a2a",
		"--a2a_agent_url", fmt.Sprintf("http://localhost:%d", port),
	})

	// 5. 런처 구성
	// 단일 에이전트 로더와 인메모리 세션 저장소를 설정합니다.
	config := &launcher.Config{
		AgentLoader:    agent.NewSingleLoader(mathAgent), // 위에서 만든 primeAgent 하나만 로드
		SessionService: session.InMemoryService(),        // 세션 데이터를 메모리에 저장 (재시작 시 초기화됨)
	}

	log.Printf("Starting Prime Server on port %d...", port)

	// 6. 서버 실행
	// 설정된 내용으로 웹 서버를 시작하고 요청을 대기합니다.
	webLauncher.Run(ctx, config)
}
