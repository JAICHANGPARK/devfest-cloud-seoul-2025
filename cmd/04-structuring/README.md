# Go ADK 핸즈온 Session 4: 정형화된 출력 (Structured Output) 만들기

환영합니다! 👋 네 번째 세션입니다.

지금까지 만든 에이전트들은 사람처럼 자연스러운 문장으로 대답했습니다. 하지만 개발자 입장에서는 에이전트의 답변을 **데이터베이스에 저장**하거나, **프론트엔드 UI에 표시**하기 위해 깔끔한 JSON 형식이 필요할 때가 많습니다.

이번 시간에는 ADK와 Gemini의 **Output Schema** 기능을 사용하여, 에이전트가 무조건 **요약(Summary)**과 **실행 항목(Action Items)** 형태의 JSON으로만 답하도록 만들어 보겠습니다.

## 🎯 학습 목표
*   **Output Schema**의 개념 이해하기
*   `genai.Schema`를 사용하여 원하는 JSON 구조 정의하기
*   비정형 텍스트(회의록, 대화 등)를 정형 데이터(JSON)로 변환하기

---

## 💻 코드 상세 분석

이번 코드의 핵심은 에이전트에게 "말하는 법"이 아니라 "답변의 틀(Format)"을 지정해 주는 것입니다.

### 1. 출력 스키마 정의 (The Blueprint) ⭐
가장 중요한 부분입니다. 에이전트가 뱉어내야 할 데이터의 구조를 정의합니다.

```go
	// 출력 스키마 정의: JSON Object 형태
	outputSchema := &genai.Schema{
		Type: genai.TypeObject, // 전체 타입은 객체(Object)
		Properties: map[string]*genai.Schema{
			// 1. summary 필드: 문자열
			"summary":      {Type: genai.TypeString},
			// 2. action_items 필드: 문자열 배열(List of Strings)
			"action_items": {Type: genai.TypeArray, Items: &genai.Schema{Type: genai.TypeString}},
		},
	}
```
*   **`genai.Schema`**: OpenAPI 스펙과 유사한 형태로 데이터 구조를 정의합니다.
*   이 설정은 모델에게 다음과 같은 제약 조건을 겁니다: *"너는 무조건 `summary`(문자열)와 `action_items`(리스트)를 가진 JSON으로만 대답해야 해."*

### 2. 에이전트에 스키마 적용
정의한 스키마를 에이전트 설정에 주입합니다.

```go
	routerAgent, err := llmagent.New(llmagent.Config{
		Name:        "root_agent", // 변수명은 routerAgent지만, 여기서는 구조화된 응답 생성기로 동작합니다.
		Model:       model,
		Description: "A helpful agent. Uses a router to route the user questions.",
		Instruction: "You are a helpful assistant. Answer the user's questions.", // 평범한 지시사항
		
		// [핵심] 출력 스키마 연결
		OutputSchema: outputSchema,
	})
```
*   **`OutputSchema`**: 이 필드가 설정되면, Gemini 모델은 `Instruction`에 있는 내용대로 생각하되, 최종 답변은 지정된 JSON 형식에 맞춰 생성합니다.

---

## 🚀 실행 및 테스트 (Let's Run!)

이 에이전트는 긴 텍스트를 정리할 때 진가를 발휘합니다. 회의록이나 긴 대화 내용을 입력으로 주어봅시다.

### 1. 실행 명령어
```bash
go run main.go run "오늘 회의에서는 다음달 마케팅 전략을 논의했어. 철수가 SNS 광고 시안을 다음 주까지 만들기로 했고, 영희는 예산안을 내일까지 정리해서 보고하기로 했어. 전체적으로 긍정적인 분위기였어."
```

### 2. 예상 실행 결과 (JSON Output)
에이전트는 더 이상 "네, 알겠습니다."라고 답하지 않습니다. 정확히 JSON 포맷으로 결과를 출력합니다.

```json
{
  "summary": "다음 달 마케팅 전략 논의 회의가 긍정적인 분위기 속에서 진행되었습니다.",
  "action_items": [
    "철수: 다음 주까지 SNS 광고 시안 제작",
    "영희: 내일까지 예산안 정리 및 보고"
  ]
}
```

### 3. 대화 모드(Chat)에서의 활용
`go run main.go chat`으로 실행한 뒤, 다음과 같이 입력해 보세요.

**User:** "라면 끓이는 법 알려줘."

**Agent:**
```json
{
  "summary": "라면을 맛있게 끓이는 기본적인 방법입니다.",
  "action_items": [
    "냄비에 물 550ml를 넣고 끓인다.",
    "물이 끓으면 면, 분말스프, 후레이크를 넣는다.",
    "4분 30초간 더 끓인다.",
    "기호에 따라 파, 계란 등을 추가한다."
  ]
}
```

---

## 🔍 활용 사례 (Use Cases)

이 기능은 다음과 같은 상황에서 매우 유용합니다.

1.  **자동 회의록 작성기**: 음성 인식(STT) 텍스트를 넣어 요약본과 할 일 목록 추출.
2.  **데이터 추출 (Extraction)**: 이메일 내용에서 주문 내역, 날짜, 연락처만 뽑아내어 DB에 저장.
3.  **API 연동**: LLM의 응답을 받아 파싱 과정 없이 바로 프론트엔드 자바스크립트 객체로 사용.

---

## 💡 팁 (Tip)

*   **변수명 주의**: 코드 상의 변수명이 `routerAgent`로 되어 있는데, 이는 에이전트가 생성된 결과를 바탕으로 다른 로직으로 '라우팅' 할 수 있다는 의미를 내포하기도 합니다. (예: action item이 있으면 Jira로 보내기 등)
*   **스키마 준수율**: Gemini 최신 모델들은 이러한 스키마 준수율이 매우 높습니다. 별도의 파싱 로직(`json.Unmarshal` 등)을 작성하여 Go 구조체로 바로 매핑할 수 있습니다.

---
수고하셨습니다! 이제 여러분은 **LLM을 확률적인 챗봇이 아닌, 예측 가능한 데이터 생성기**로 다루는 강력한 무기를 얻었습니다. 😎