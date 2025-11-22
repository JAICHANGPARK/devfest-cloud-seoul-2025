# devfest-cloud-seoul-2025: Building AI Agents with Go ADK
# devfest-cloud-seoul-2025: Go ADK를 활용한 AI 에이전트 구축 핸즈온

## English Version

This repository contains the hands-on session materials for DevFest Cloud Seoul 2025, focusing on building AI agents with Google's **Agent Development Kit (ADK) for Go** and the Gemini model.

### Project Overview
This project guides participants through various aspects of developing AI agents, progressing from fundamental concepts to advanced topics like multi-agent workflows and distributed systems. Each module in the `cmd/` directory represents a distinct hands-on session, offering practical experience with the ADK framework.

### Hands-on Sessions

#### 01-hello-agent: Hello Agent (First AI Agent)
- **Goal**: Build a basic AI agent using Google's Agent Development Kit (ADK) for Go, based on the Gemini model. Understand ADK's core components: Model, Agent, and Launcher.
- **Key Concepts**: Gemini Model, ADK Agent, ADK Launcher, Basic Agent Configuration.

#### 02-search-tool: Search Agent (Tools for Search)
- **Goal**: Enhance the agent with Google Search capabilities using ADK's Tools system to access real-time information.
- **Key Concepts**: ADK Tool Interface (`google.golang.org/adk/tool`), Google Search Integration (`geminitool.GoogleSearch{}`), Instruction Tuning, Grounding.

#### 03-custom-tools: Custom Tools (Creating Your Own Custom Tools)
- **Goal**: Implement Function Calling, allowing the agent to execute custom Go functions for tasks like querying a database, calling internal APIs, or running specific business logic.
- **Key Concepts**: Custom Tool Structure, `functiontool` Package, `jsonschema` Tags for argument information, Tool Chaining, `Description` Field importance.

#### 04-structuring: Structured Output (Creating Structured Output)
- **Goal**: Make the agent respond in a structured JSON format (Summary and Action Items) using ADK and Gemini's Output Schema feature.
- **Key Concepts**: Output Schema (`genai.Schema`), Defining JSON Structures, Structured Data Generation, Use Cases (meeting minutes, data extraction, API integration).

#### 05-structuring-tuned: Intelligent Router Agent (Intelligent Router Agent)
- **Goal**: Create a Router Agent that classifies user intent and determines which agent should handle the request, rather than directly answering.
- **Key Concepts**: Classifier Pattern, Enum Schema for restricted output, Flash Model Usage for speed, Negative Constraints in instructions, Metadata in Schema.

#### 06-session-memory: Memory & Runner (Agent with Memory)
- **Goal**: Give the agent "long-term memory" capabilities using ADK's `Runner` to control conversational flow and `Memory Service` to store and retrieve past conversations.
- **Key Concepts**: `Runner` vs. `Launcher`, `Memory Service` (session & long-term memory), Memory Tool (`search_past_conversations`), Prompt Engineering for Korean search accuracy, Manual Memory Storage, RAG (Retrieval-Augmented Generation).

#### 07-trip-planner: Parallel & Sequential (Multi-Agent Workflow)
- **Goal**: Build a "Trip Planner Avengers" team by combining `Parallel Agent` and `Sequential Agent` to create a multi-agent workflow.
- **Key Concepts**: Multi-Agent Architecture, `Parallel Agent` for concurrent execution, `Sequential Agent` for pipeline construction, `OutputKey` and `{variable}` syntax for context sharing, Division of Labor.

#### 08-a2a: Agent-to-Agent (Remote Agents and A2A)
- **Goal**: Learn ADK's powerful A2A (Agent-to-Agent) protocol to enable agents to communicate with and delegate tasks to remote agents running on different servers.
- **Key Concepts**: Remote Agent Architecture (`web.Launcher`), A2A Protocol (Agent Card), Tool Expansion (multi-tool server agent), Client-Server Model, Microservices, Security, Language Agnosticism.

### Getting Started

Each session folder includes a `README.md` file with detailed instructions on how to set up the environment, run the specific examples, and explanations of the code.

To run any of the examples, navigate to the respective module directory (e.g., `cmd/01-hello-agent`) and execute:

```bash
go mod tidy
go run main.go [chat|run "Your question"]
```
*(Specific commands might vary slightly per session, refer to individual `README.md` files.)*

### Prerequisites
1.  **Go Installation**: Go version 1.21 or higher is required.
2.  **Google Cloud Project & API Key**: An API key is needed to use the Gemini API.
3.  **Environment Variable Setup**: Set your API key as the `GOOGLE_API_KEY` environment variable.
    ```bash
    export GOOGLE_API_KEY="YOUR_ACTUAL_API_KEY"
    ```

---

## 한국어 버전

이 저장소는 DevFest Cloud Seoul 2025에서 Google의 **Go용 Agent Development Kit (ADK)**와 Gemini 모델을 활용하여 AI 에이전트를 구축하는 핸즈온 세션 자료를 포함합니다.

### 프로젝트 개요
이 프로젝트는 AI 에이전트 개발의 기본 개념부터 멀티 에이전트 워크플로우 및 분산 시스템과 같은 고급 주제에 이르기까지 다양한 측면을 참가자들에게 안내합니다. `cmd/` 디렉토리의 각 모듈은 ADK 프레임워크를 통한 실습 경험을 제공하는 개별 핸즈온 세션을 나타냅니다.

### 핸즈온 세션

#### 01-hello-agent: 첫 번째 AI 에이전트 만들기
- **목표**: Gemini 모델 기반의 Google Agent Development Kit (ADK) for Go를 사용하여 기본적인 AI 에이전트를 구축합니다. ADK의 핵심 구성 요소인 Model, Agent, Launcher의 상호작용을 이해합니다.
- **핵심 개념**: Gemini 모델, ADK 에이전트, ADK 런처, 기본적인 에이전트 구성.

#### 02-search-tool: 도구(Tools)를 활용한 검색 에이전트 만들기
- **목표**: ADK의 Tools 시스템을 사용하여 에이전트에 Google Search 기능을 부여하고 실시간 정보에 접근하도록 합니다.
- **핵심 개념**: ADK Tool 인터페이스 (`google.golang.org/adk/tool`), Google 검색 연동 (`geminitool.GoogleSearch{}`), Instruction 튜닝, 그라운딩(Grounding).

#### 03-custom-tools: 나만의 커스텀 도구(Custom Tools) 만들기
- **목표**: Function Calling을 구현하여 에이전트가 데이터베이스 조회, 내부 API 호출 또는 특정 비즈니스 로직 실행과 같은 작업을 위해 사용자 지정 Go 함수를 실행할 수 있도록 합니다.
- **핵심 개념**: 커스텀 도구 구조, `functiontool` 패키지, 인자 정보를 위한 `jsonschema` 태그, 도구 체이닝, `Description` 필드의 중요성.

#### 04-structuring: 정형화된 출력 (Structured Output) 만들기
- **목표**: ADK와 Gemini의 Output Schema 기능을 사용하여 에이전트가 요약(Summary) 및 실행 항목(Action Items) 형태의 정형화된 JSON으로 응답하도록 만듭니다.
- **핵심 개념**: Output Schema (`genai.Schema`), JSON 구조 정의, 정형화된 데이터 생성, 활용 사례 (회의록, 데이터 추출, API 연동).

#### 05-structuring-tuned: 지능형 라우터 에이전트 (Intelligent Router Agent)
- **목표**: 사용자 의도를 분류하고 직접 답변하는 대신 어떤 에이전트가 요청을 처리해야 하는지 결정하는 라우터 에이전트를 만듭니다.
- **핵심 개념**: 분류기 패턴, 열거형(Enum) 스키마를 통한 출력 제한, Flash 모델 활용 (속도 최적화), Instruction의 부정적 제약 조건, 스키마 내 메타데이터.

#### 06-session-memory: 기억을 가진 에이전트 (Memory & Runner)
- **목표**: ADK의 `Runner`를 사용하여 대화 흐름을 제어하고 `Memory Service`를 통해 에이전트에 "장기 기억" 기능을 부여하여 과거 대화를 저장하고 검색합니다.
- **핵심 개념**: `Runner` vs. `Launcher`, `Memory Service` (세션 및 장기 기억), 메모리 도구 (`search_past_conversations`), 한국어 검색 정확도를 위한 프롬프트 엔지니어링, 수동 메모리 저장, RAG (검색 증강 생성).

#### 07-trip-planner: 멀티 에이전트 워크플로우 (Parallel & Sequential)
- **목표**: `Parallel Agent`와 `Sequential Agent`를 결합하여 멀티 에이전트 워크플로우인 "여행 계획 어벤져스" 팀을 구축합니다.
- **핵심 개념**: 멀티 에이전트 아키텍처, 병렬 실행을 위한 `Parallel Agent`, 파이프라인 구성을 위한 `Sequential Agent`, `OutputKey` 및 `{variable}` 문법을 통한 컨텍스트 공유, 역할 분담.

#### 08-a2a: 원격 에이전트와 A2A (Agent-to-Agent)
- **목표**: ADK의 강력한 A2A (Agent-to-Agent) 프로토콜을 학습하여 에이전트가 다른 서버에서 실행되는 원격 에이전트와 통신하고 작업을 위임할 수 있도록 합니다.
- **핵심 개념**: 원격 에이전트 아키텍처 (`web.Launcher`), A2A 프로토콜 (Agent Card), 도구 확장 (다중 도구 서버 에이전트), 클라이언트-서버 모델, 마이크로서비스, 보안, 언어 독립성.

### 시작하기

각 세션 폴더에는 환경 설정, 특정 예제 실행 방법 및 코드 설명에 대한 자세한 지침이 담긴 `README.md` 파일이 있습니다.

예제를 실행하려면 해당 모듈 디렉토리(예: `cmd/01-hello-agent`)로 이동하여 다음을 실행합니다.

```bash
go mod tidy
go run main.go [chat|run "당신의 질문"]
```
*(세션별로 명령이 약간 다를 수 있으므로 개별 `README.md` 파일을 참조하십시오.)*

### 사전 준비 사항
1.  **Go 설치**: Go 1.21 이상 버전이 필요합니다.
2.  **Google Cloud Project & API Key**: Gemini API를 사용하기 위한 API 키가 필요합니다.
3.  **환경 변수 설정**: API 키를 `GOOGLE_API_KEY` 환경 변수로 설정해야 합니다.
    ```bash
    export GOOGLE_API_KEY="YOUR_ACTUAL_API_KEY"
    ```