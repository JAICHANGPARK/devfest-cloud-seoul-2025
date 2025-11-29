package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	// 라우팅은 속도가 생명이므로 Flash 모델 권장 (예: gemini-1.5-flash)
	model, err := gemini.NewModel(ctx,
		"gemini-3-pro-preview",
		&genai.ClientConfig{
			APIKey: os.Getenv("GOOGLE_API_KEY"),
		})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	// [개선 1] OutputSchema: 답변이 아닌 '라우팅 결정'을 위한 구조체 정의
	// 사용자의 의도를 파악하여 정해진 카테고리 중 하나로 분류합니다.
	outputSchema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			// 1. 어디로 보낼지 결정 (Enum을 사용하여 환각 방지 및 엄격한 분류)
			"destination": {
				Type:        genai.TypeString,
				Enum:        []string{"technical_support", "billing_inquiry", "general_chat", "escalate_to_human"},
				Description: "The target agent or department to handle the user query.",
			},
			// 2. 분류 이유 (디버깅 및 검증용)
			"reasoning": {
				Type:        genai.TypeString,
				Description: "Explanation of why this destination was chosen.",
			},
			// 3. 사용자 의도 요약 (다음 에이전트에게 넘겨주기 위함)
			"intent_summary": {
				Type:        genai.TypeString,
				Description: "A concise summary of what the user wants to achieve.",
			},
			// 4. 난이도/우선순위 파악
			"priority": {
				Type: genai.TypeString,
				Enum: []string{"high", "medium", "low"},
			},
		},
		// 필수 필드 지정
		Required: []string{"destination", "reasoning", "intent_summary"},
	}

	// [개선 2] Instruction: 역할을 '분류자(Classifier)'로 명확히 정의
	instruction := `
You are an intelligent request router. 
Your goal is NOT to answer the user's question directly, but to classify the intent and route it to the correct department.

Classify the input into one of the following destinations:
1. 'technical_support': Questions about code, bugs, installation, or technical errors.
2. 'billing_inquiry': Questions about payments, invoices, pricing, or subscriptions.
3. 'general_chat': Greetings, small talk, or non-specific questions.
4. 'escalate_to_human': Complex complaints, legal issues, or when the user is very angry.

Analyze the user's input carefully and determine the destination, priority, and a summary of their intent.
`

	routerAgent, err := llmagent.New(llmagent.Config{
		Name:         "router_agent", // 이름도 역할에 맞게 변경
		Model:        model,
		Description:  "Analyzes user input and routes it to the appropriate specialized agent.",
		Instruction:  instruction,
		OutputSchema: outputSchema,
	})

	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(routerAgent),
	}

	l := full.NewLauncher()

	// 실행 시 인자 예시: "내 신용카드 결제가 두 번 되었어, 환불해줘" -> billing_inquiry
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
