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
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx,
		"gemini-2.5-flash-lite",
		//"gemini-3-pro-preview",
		&genai.ClientConfig{
			APIKey: os.Getenv("GOOGLE_API_KEY"),
		})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	timeAgent, err := llmagent.New(llmagent.Config{
		Name:        "search_agent",
		Model:       model,
		Description: "A helpful agent that searches the web.",
		Instruction: "You are a helpful assistant. Use Google Search to answer the user's questions.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
	})

	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(timeAgent),
	}

	l := full.NewLauncher()

	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
