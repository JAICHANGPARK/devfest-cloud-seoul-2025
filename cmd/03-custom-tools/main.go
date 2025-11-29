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
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/functiontool"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx,
		"gemini-3-pro-preview",
		//"gemini-3-pro-preview",
		&genai.ClientConfig{
			APIKey: os.Getenv("GOOGLE_API_KEY"),
		})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	weatherTool, _ := functiontool.New(functiontool.Config{
		Name: "get_weather", Description: "Get weather for a city"},
		getWeather,
	)

	sentimentTool, _ := functiontool.New(
		functiontool.Config{Name: "analyze_sentiment", Description: "Analyze text sentiment"},
		analyzeSentiment)

	myAgent, err := llmagent.New(llmagent.Config{
		Name:  "helper_agent",
		Model: model,
		Instruction: "You are a helper. If asked about weather, use get_weather.  " +
			"Then analyze the user's reaction using analyze_sentiment.",
		Tools: []tool.Tool{weatherTool, sentimentTool},
	})

	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(myAgent),
	}

	l := full.NewLauncher()

	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
