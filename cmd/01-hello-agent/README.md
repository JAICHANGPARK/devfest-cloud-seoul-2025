# Go ADK (Agent Development Kit) 핸즈온: 첫 번째 AI 에이전트 만들기

환영합니다! 👋 이번 핸즈온 세션에서는 Google의 **Agent Development Kit (ADK) for Go**를 사용하여 Gemini 모델 기반의 AI 에이전트를 만들어 보겠습니다.

이 가이드는 제공된 예제 코드를 단계별로 분석하여, ADK의 핵심 구성 요소인 **Model**, **Agent**, **Launcher**가 어떻게 상호작용하는지 이해하는 것을 목표로 합니다.

### 📋 개요
우리가 만들 프로그램은 사용자의 질문에 답변하는 기본적인 "Helpful Assistant"입니다. ADK 프레임워크를 사용하면 복잡한 LLM 연동 로직을 표준화된 방식으로 구현할 수 있습니다.

### 🛠️ 사전 준비 사항 (Prerequisites)
1.  **Go 설치**: Go 1.21 이상 버전이 필요합니다.
2.  **Google Cloud Project & API Key**: Gemini API를 사용하기 위한 API 키가 필요합니다.
3.  **환경 변수 설정**: API 키를 `GOOGLE_API_KEY` 환경 변수로 설정해야 합니다.

```bash
export GOOGLE_API_KEY="YOUR_ACTUAL_API_KEY"
```

---

### 💻 코드 상세 분석 (`main.go`)

작성된 `main.go` 코드를 논리적인 블록으로 나누어 살펴보겠습니다.

-   **`package main`**: 이 코드가 실행 가능한 프로그램임을 나타내는 `main` 패키지를 선언합니다.
-   **임포트(Imports)**:
    -   `context`: 요청 수명 주기, 취소 및 마감일을 관리합니다.
    -   `log`: 특히 치명적인 오류 메시지를 로깅하는 데 사용됩니다.
    -   `os`: 환경 변수(예: API 키) 및 명령줄 인수에 액세스하는 데 사용됩니다.
    -   `google.golang.org/adk/...`: 이 임포트들은 Gemini Agent Development Kit (ADK)에서 가져온 것입니다. 핵심 에이전트 기능, LLM 특정 에이전트 및 런처를 포함하여 AI 에이전트를 구축하기 위한 프레임워크를 제공합니다.
    -   `google.golang.org/genai`: Gemini 모델에 대한 연결을 구성하는 데 사용되는 Google AI Go SDK입니다.
-   **`func main()`**: 프로그램의 진입점입니다.
    -   `ctx := context.Background()`: Go에서 작업 관리를 위한 표준 관행인 백그라운드 컨텍스트를 초기화합니다.
    -   `model, err := gemini.NewModel(...)`: 이 줄은 Gemini 언어 모델 인스턴스를 생성합니다. `"gemini-2.5-flash-lite"`를 모델 이름으로 지정하고 `GOOGLE_API_KEY` 환경 변수에서 가져온 API 키로 클라이언트를 구성합니다. 모델 생성에 대한 오류 처리도 포함되어 있습니다.
    -   `rootAgent, err := llmagent.New(...)`: 핵심 AI 에이전트를 생성합니다. 이름(`"root_agent"`)이 부여되고 이전에 생성된 `model`에 연결되며, 동작을 정의하는 `Description` 및 `Instruction` ("당신은 유용한 비서입니다. 사용자의 질문에 답변하세요.")이 제공됩니다.
    -   `config := &launcher.Config{...}`: 에이전트 런처 구성을 설정합니다. `agent.NewSingleLoader(rootAgent)`는 애플리케이션 시작 시 `rootAgent`가 로드되도록 합니다.
    -   `l := full.NewLauncher()`: 에이전트의 수명 주기를 관리할 전체 런처를 초기화합니다.
    -   `if err = l.Execute(...)`: 에이전트가 실행되는 곳입니다. 컨텍스트, 런처 구성 및 스크립트에 전달된 명령줄 인수(`os.Args[1:]`)를 사용합니다. 실행 중에 오류가 발생하면 오류를 로깅하고 올바른 명령줄 구문을 표시합니다.

### 🚀 실행 방법 (How to Run)
코드를 작성한 후 터미널에서 아래와 같이 실행해 보세요.

1.  **의존성 설치**
    ```bash
    go mod tidy
    ```

2.  **대화형 모드(Chat)로 실행**
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

3.  **단발성 질문 실행**
    ```bash
    go run main.go run "Go 언어의 장점을 한 문장으로 설명해줘"
    ```

---

### 💡 팁 & 트러블슈팅

*   **403 Permission Denied**: `GOOGLE_API_KEY`가 올바르게 설정되었는지, 해당 키가 Gemini API를 사용할 권한이 있는지 확인하세요.
*   **Model Not Found**: 코드에 적힌 모델명(`gemini-2.5-flash-lite`)이 현재 사용 가능한지 확인하세요. 만약 오류가 난다면 `gemini-2.5-flash`로 변경해 보세요.
*   **프롬프트 수정**: `Instruction` 필드의 내용을 바꿔보세요. (예: "You are a pirate."라고 입력하면 해적 말투로 대답합니다.)

---

Happy Coding! 🎉 ADK로 나만의 멋진 AI 에이전트를 확장해 보세요.

--- 

# Hello Agent (첫 번째 AI 에이전트 만들기)

## English Version

Welcome! 👋 In this hands-on session, we will build an AI agent based on the Gemini model using Google's **Agent Development Kit (ADK) for Go**.

This guide aims to analyze the provided example code step-by-step to understand how the core components of ADK – **Model**, **Agent**, and **Launcher** – interact.

### 📋 Overview
Our program will be a basic "Helpful Assistant" that answers user questions. The ADK framework allows us to implement complex LLM integration logic in a standardized way.

### 🛠️ Prerequisites
1.  **Go Installation**: Go version 1.21 or higher is required.
2.  **Google Cloud Project & API Key**: An API key is needed to use the Gemini API.
3.  **Environment Variable Setup**: Set your API key as the `GOOGLE_API_KEY` environment variable.

```bash
export GOOGLE_API_KEY="YOUR_ACTUAL_API_KEY"
```

### 💻 Code Explanation (`main.go`)

This Go program sets up and runs a basic AI agent using the Gemini ADK. Here's a breakdown of the key components:

-   **`package main`**: Declares the package as `main`, indicating that this code is an executable program.
-   **Imports**:
    -   `context`: For managing request lifecycles, cancellations, and deadlines.
    -   `log`: For logging messages, particularly fatal errors.
    -   `os`: Used here to access environment variables (like API keys) and command-line arguments.
    -   `google.golang.org/adk/...`: These imports are from the Gemini Agent Development Kit (ADK). They provide the framework for building AI agents, including core agent functionalities, LLM-specific agents, and a launcher to run them.
    -   `google.golang.org/genai`: The Google AI Go SDK, used here to configure the connection to the Gemini model.
-   **`func main()`**: The entry point of the program.
    -   `ctx := context.Background()`: Initializes a background context, which is standard practice in Go for managing operations.
    -   `model, err := gemini.NewModel(...)`: This line creates an instance of the Gemini language model. It specifies `"gemini-2.5-flash-lite"` as the model name and configures the client with an API key fetched from the `GOOGLE_API_KEY` environment variable. Error handling is included for model creation.
    -   `rootAgent, err := llmagent.New(...)`: This creates the core AI agent. It's given a name (`"root_agent"`), linked to the `model` created earlier, and provided with a `Description` and `Instruction` that define its behavior ("You are a helpful assistant. Answer the user's questions.").
    -   `config := &launcher.Config{...}`: Sets up the configuration for the agent launcher. `agent.NewSingleLoader(rootAgent)` ensures that our `rootAgent` is loaded when the application starts.
    -   `l := full.NewLauncher()`: Initializes the full launcher, which will manage the agent's lifecycle.
    -   `if err = l.Execute(...)`: This is where the agent is executed. It takes the context, the launcher configuration, and any command-line arguments passed to the script (`os.Args[1:]`). If any error occurs during execution, it logs the error and shows the correct command-line syntax.

### 🚀 How to Run
After writing the code, run it in the terminal as follows:

1.  **Install Dependencies**
    ```bash
    go mod tidy
    ```

2.  **Run in Chat Mode**
    Thanks to the ADK Launcher, you can use chat mode directly without additional implementation.
    ```bash
    go run main.go chat
    ```
    **Example Output:**
    ```text
    Type "exit" or "quit" to stop the session.
    >>> Hello!
    Hello! How can I help you today?
    ```

3.  **Run with a Single Question**
    ```bash
    go run main.go run "Describe the advantages of Go language in one sentence"
    ```

### 💡 Tips & Troubleshooting
*   **403 Permission Denied**: Check if `GOOGLE_API_KEY` is correctly set and has permission to use the Gemini API.
*   **Model Not Found**: Verify that the model name specified in the code (`gemini-2.5-flash-lite`) is currently available. If an error occurs, try changing it to `gemini-2.5-flash`.
*   **Modify Prompt**: Try changing the content of the `Instruction` field (e.g., if you enter "You are a pirate.", it will respond in a pirate accent).

