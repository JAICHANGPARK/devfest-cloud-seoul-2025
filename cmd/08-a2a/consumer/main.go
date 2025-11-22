package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/remoteagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	// 1. 모델 설정
	model, err := gemini.NewModel(ctx,
		"gemini-2.5-flash-lite",
		&genai.ClientConfig{
			APIKey: os.Getenv("GOOGLE_API_KEY"),
		})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	sessionService := session.InMemoryService()

	// 2. 원격 에이전트(A2A) 정의 수정
	// 이전에는 "RemotePrimeAgent"였으나, 이제는 수학 전반을 다루므로 이름을 변경하고
	// Description에 새로운 능력(팩토리얼, GCD)을 명시해야 합니다.
	remoteMathAgent, err := remoteagent.NewA2A(remoteagent.A2AConfig{
		Name: "RemoteMathHelper", // 이름 변경
		// [중요] 이 설명(Description)을 보고 메인 에이전트가 작업을 위임할지 결정합니다.
		Description:     "Can check prime numbers, calculate factorials, and find GCD of two numbers.",
		AgentCardSource: "http://localhost:8001", // 서버 주소 (앞서 만든 서버가 8001 포트)
	})
	if err != nil {
		log.Fatal(err)
	}

	// 3. 메인 에이전트(MathTutor) 수정
	mathTutor, _ := llmagent.New(llmagent.Config{
		Name:  "MathTutor",
		Model: model,
		// 지시문 수정: 소수뿐만 아니라 다른 수학 질문도 원격 에이전트에게 물어보라고 지시
		Instruction: "You are a math tutor. If the user asks about checking primes, calculating factorials, or finding the GCD, delegate the task to the RemoteMathHelper.",
		// SubAgents에 원격 에이전트 등록
		SubAgents: []agent.Agent{remoteMathAgent},
	})

	// 4. 런처 실행 설정
	config := &launcher.Config{
		AgentLoader:    agent.NewSingleLoader(mathTutor),
		SessionService: sessionService,
	}

	// CLI(터미널) 모드로 실행
	l := full.NewLauncher()

	// 실행 (터미널에서 질문 입력 가능)
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
