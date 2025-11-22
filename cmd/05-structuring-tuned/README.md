# Go ADK 핸즈온 Session 5: 지능형 라우터 에이전트 (Intelligent Router Agent)

환영합니다! 👋 다섯 번째 세션입니다.

지금까지 우리는 하나의 에이전트가 도구도 쓰고, 답변도 하는 방식을 배웠습니다. 하지만 시스템이 커지면 하나의 에이전트가 모든 것을 처리할 수 없습니다. 개발 전문 에이전트, 결제 전문 에이전트, 상담원 연결 등이 필요하죠.

이번 시간에는 사용자의 말을 듣고 **"직접 답변하지 않고, 누가 처리해야 할지 결정하는"** 라우터(Router) 에이전트를 만들어 봅니다.

## 🎯 학습 목표
*   **Classifier Pattern**: LLM을 생성기가 아닌 '분류기'로 사용하는 패턴 이해하기
*   **Enum Schema**: 출력 값을 특정 키워드로 제한하여 프로그램 제어력 높이기
*   **Flash Model**: 단순/반복 작업에 최적화된 빠르고 가벼운 모델(Flash)의 적재적소 활용

---

## 💻 코드 상세 분석

이번 코드는 에이전트에게 "말하지 말고, 판단하라"고 시키는 것이 핵심입니다.

### 1. 모델 선정 (Speed is Key) ⚡
```go
	// 라우팅은 속도가 생명! Gemini Flash 모델 사용
	model, err := gemini.NewModel(ctx, "gemini-2.5-flash", ...)
```
*   라우터는 사용자와의 대화 첫 관문입니다. 여기서 시간이 지체되면 전체 응답 속도가 느려집니다.
*   따라서 가장 성능이 뛰어나고 비용이 저렴한 **Flash** 계열 모델이 라우팅 작업에 가장 적합합니다.

### 2. 라우팅 스키마 정의 (The Logic) ⭐
에이전트가 내릴 수 있는 결정의 범위를 코드로 강제합니다.

```go
	outputSchema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			// 1. 목적지 (Destination) - Enum 활용
			"destination": {
				Type: genai.TypeString,
				// [중요] 오타나 엉뚱한 단어가 나오지 않도록 선택지를 고정합니다.
				Enum: []string{"technical_support", "billing_inquiry", "general_chat", "escalate_to_human"},
			},
			// 2. 판단 근거 (Reasoning)
			"reasoning": { Type: genai.TypeString },
			// 3. 우선순위 (Priority)
			"priority": {
				Type: genai.TypeString,
				Enum: []string{"high", "medium", "low"},
			},
		},
        // ...
	}
```
*   **Enum (열거형)**: LLM은 창의적이라 때로는 "billing"을 "finance"나 "money_help"라고 맘대로 바꿀 수 있습니다. `Enum`을 사용하면 코드에서 `if destination == "billing_inquiry"` 처럼 안전하게 분기 처리를 할 수 있습니다.
*   **Reasoning**: 에이전트가 왜 그런 판단을 했는지 로그를 남겨 디버깅할 수 있게 합니다.

### 3. 지시사항 (The Role)
```go
	instruction := `
    You are an intelligent request router. 
    Your goal is NOT to answer the user's question directly, 
    but to classify the intent and route it to the correct department.
    ...
    `
```
*   **Negative Constraint**: "직접 답변하지 말라(NOT to answer)"는 제약을 걸어, 라우터 본연의 임무에 집중하게 합니다.

---

## 🚀 실행 및 테스트 (Let's Run!)

다양한 상황을 연출하여 라우터가 올바르게 분류하는지 확인해 봅시다.

### 1. 기술 지원 요청 (Technical Support)
```bash
go run main.go run "서버 로그에 500 에러가 계속 뜨고 배포가 안 돼요. 급합니다!"
```
**예상 결과:**
```json
{
  "destination": "technical_support",
  "priority": "high",
  "reasoning": "User is reporting a server error (500) and deployment failure, indicating a technical issue.",
  "intent_summary": "Deployment failure with 500 error logs."
}
```
*   `technical_support`로 분류되었고, "급합니다"라는 말과 에러 상황을 보고 `priority`를 `high`로 잡았습니다.

### 2. 환불/결제 문의 (Billing)
```bash
go run main.go run "지난달 요금이 왜 이렇게 많이 나왔죠? 확인 부탁드립니다."
```
**예상 결과:**
```json
{
  "destination": "billing_inquiry",
  "priority": "medium",
  "reasoning": "User is asking about an unexpectedly high bill charge.",
  "intent_summary": "Inquiry about high billing amount for last month."
}
```

### 3. 상담원 연결 (Escalation - 감정 분석 포함)
```bash
go run main.go run "아니 상담원 연결해달라고 몇 번을 말해! 지금 장난해?"
```
**예상 결과:**
```json
{
  "destination": "escalate_to_human",
  "priority": "high",
  "reasoning": "User is expressing anger and explicitly demanding a human agent.",
  "intent_summary": "Angry user demanding human intervention."
}
```
*   단순 키워드 매칭이 아니라, 문맥 속의 **분노(Anger)**를 감지하여 `escalate_to_human`으로 보냅니다.

---

## 🔍 활용 방안 (Next Steps)

이 라우터 에이전트는 실제 시스템에서 다음과 같이 활용됩니다.

1.  **Switch 문 구현**: Go 코드에서 `destination` 값에 따라 다른 함수나 API를 호출합니다.
    ```go
    // 예시 의사 코드 (Pseudo-code)
    resp := routerAgent.Run(input)
    switch resp.Destination {
    case "technical_support":
        jiraAgent.CreateTicket(resp.IntentSummary)
    case "billing_inquiry":
        billingTool.CheckStatus(userID)
    case "escalate_to_human":
        slack.Alert("Angry customer detected!")
    }
    ```
2.  **비용 절감**: 모든 질문을 비싼 고성능 모델(Pro/Ultra)로 처리하는 대신, 앞단의 가벼운 라우터가 분류하여 간단한 인사는 무시하거나 저렴한 모델로 넘길 수 있습니다.

---

## 💡 팁 (Tip)

*   **스키마의 힘**: `OutputSchema`에 `priority` 같은 메타데이터 필드를 추가하면, 단순 분류를 넘어 업무의 경중까지 판단해 주는 아주 똑똑한 비서가 됩니다.
*   **프롬프트 튜닝**: 분류가 잘 안 된다면 `Instruction` 부분에 예시(Few-shot prompting)를 몇 개 추가해 주면 정확도가 비약적으로 상승합니다.

---
수고하셨습니다! 여러분은 이제 AI 시스템의 **교통 정리를 담당하는 관제탑**을 건설했습니다. 🚦🛫