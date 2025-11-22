# Go ADK (Agent Development Kit) 핸즈온: 첫 번째 AI 에이전트 만들기

환영합니다! 👋 이번 핸즈온 세션에서는 Google의 **Agent Development Kit (ADK) for Go**를 사용하여 Gemini 모델 기반의 AI 에이전트를 만들어 보겠습니다.

이 가이드는 제공된 예제 코드를 단계별로 분석하여, ADK의 핵심 구성 요소인 **Model**, **Agent**, **Launcher**가 어떻게 상호작용하는지 이해하는 것을 목표로 합니다.

## 📋 개요
우리가 만들 프로그램은 사용자의 질문에 답변하는 기본적인 "Helpful Assistant"입니다. ADK 프레임워크를 사용하면 복잡한 LLM 연동 로직을 표준화된 방식으로 구현할 수 있습니다.

## 🛠️ 사전 준비 사항 (Prerequisites)
1. **Go 설치**: Go 1.21 이상 버전이 필요합니다.
2. **Google Cloud Project & API Key**: Gemini API를 사용하기 위한 API 키가 필요합니다.
3. **환경 변수 설정**: API 키를 `GOOGLE_API_KEY` 환경 변수로 설정해야 합니다.

```bash
export GOOGLE_API_KEY="YOUR_ACTUAL_API_KEY"
```

---

## 💻 코드 상세 분석

작성된 `main.go` 코드를 논리적인 블록으로 나누어 살펴보겠습니다.

### 1. 패키지 및 라이브러리 임포트
가장 먼저 필요한 ADK 라이브러리들을 가져옵니다.

```go
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
```
*   `google.golang.org/adk/...`: 에이전트 개발을 위한 핵심 프레임워크입니다.
*   `google.golang.org/genai`: Google GenAI(Gemini) Go 클라이언트 SDK입니다.

### 2. 모델(Model) 초기화
에이전트의 "두뇌" 역할을 할 Gemini 모델을 설정합니다.

```go
func main() {
	ctx := context.Background()

	// Gemini 모델 클라이언트 생성
	model, err := gemini.NewModel(ctx,
		"gemini-2.5-flash-lite", // 사용할 모델 버전 (가볍고 빠른 모델 선택)
		&genai.ClientConfig{
			APIKey: os.Getenv("GOOGLE_API_KEY"), // 환경 변수에서 API 키 로드
		})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}
    // ...
```
*   **`gemini.NewModel`**: ADK 호환 모델 객체를 생성합니다.
*   **모델 선택**: 여기서는 `gemini-2.5-flash-lite`를 사용합니다. 필요에 따라 `gemini-1.5-pro` 등으로 변경 가능합니다.
*   **APIKey**: 코드에 직접 키를 하드코딩하지 않고 `os.Getenv`를 통해 보안을 유지합니다.

### 3. 에이전트(Agent) 정의
모델을 사용하여 실제로 동작할 "에이전트"를 정의합니다.

```go
    // ...
	rootAgent, err := llmagent.New(llmagent.Config{
		Name:        "root_agent",    // 에이전트의 고유 식별자
		Model:       model,           // 위에서 생성한 Gemini 모델 연결
		Description: "A helpful agent.", // 이 에이전트가 무엇을 하는지 설명
		Instruction: "You are a helpful assistant. Answer the user's questions.", // 시스템 프롬프트 (페르소나 설정)
	})

	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}
    // ...
```
*   **`llmagent.New`**: LLM 기반 에이전트를 생성합니다.
*   **`Instruction`**: 가장 중요한 부분입니다. 에이전트에게 "너는 누구이고, 어떻게 행동해야 해"라고 지시하는 시스템 프롬프트 역할을 합니다.

### 4. 런처(Launcher) 구성 및 실행
생성한 에이전트를 실행 환경(CLI 등)과 연결하는 단계입니다.

```go
    // ...
    // 런처 설정: 단일 에이전트 로더 사용
	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(rootAgent),
	}

    // 전체 기능을 갖춘 런처 생성
	l := full.NewLauncher()

    // 런처 실행: CLI 인자(os.Args)를 받아 처리
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
```
*   **`agent.NewSingleLoader`**: 복잡한 에이전트 라우팅 없이, 우리가 만든 `rootAgent` 하나만 실행하도록 설정합니다.
*   **`full.NewLauncher()`**: ADK가 제공하는 표준 실행기입니다. 이를 사용하면 별도의 채팅 루프를 짤 필요 없이, 즉시 CLI 명령어로 에이전트와 대화할 수 있습니다.
*   **`l.Execute`**: 프로그램을 실행합니다. 사용자가 터미널에 입력한 명령어(예: `chat`)를 해석하여 에이전트를 구동합니다.

---

## 🚀 실행 방법 (How to Run)

코드를 작성한 후 터미널에서 아래와 같이 실행해 보세요.

### 1. 의존성 설치
```bash
go mod tidy
```

### 2. 대화형 모드(Chat)로 실행
ADK Launcher 덕분에 별도 구현 없이 바로 채팅 모드를 사용할 수 있습니다.

```bash
go run main.go chat
```

**실행 결과 예시:**
```text
Type "exit" or "quit" to stop the session.
>>> 안녕하세요!
Hello! How can I help you today?
```

### 3. 단발성 질문 실행
```bash
go run main.go run "Go 언어의 장점을 한 문장으로 설명해줘"
```

---

## 💡 팁 & 트러블슈팅

*   **403 Permission Denied**: `GOOGLE_API_KEY`가 올바르게 설정되었는지, 해당 키가 Gemini API를 사용할 권한이 있는지 확인하세요.
*   **Model Not Found**: 코드에 적힌 모델명(`gemini-2.5-flash-lite`)이 현재 사용 가능한지 확인하세요. 만약 오류가 난다면 `gemini-2.5-flash`로 변경해 보세요.
*   **프롬프트 수정**: `Instruction` 필드의 내용을 바꿔보세요. (예: "You are a pirate."라고 입력하면 해적 말투로 대답합니다.)

---
Happy Coding! 🎉 ADK로 나만의 멋진 AI 에이전트를 확장해 보세요.