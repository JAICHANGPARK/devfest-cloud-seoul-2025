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

	model, err := gemini.NewModel(ctx,
		"gemini-2.5-flash-lite",
		//"gemini-3-pro-preview",
		&genai.ClientConfig{
			APIKey: os.Getenv("GOOGLE_API_KEY"),
		})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	outputSchema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"summary":      {Type: genai.TypeString},
			"action_items": {Type: genai.TypeArray, Items: &genai.Schema{Type: genai.TypeString}},
		},
	}

	routerAgent, err := llmagent.New(llmagent.Config{
		Name:         "root_agent",
		Model:        model,
		Description:  "A helpful agent. Uses a router to route the user questions.",
		Instruction:  "You are a helpful assistant. Answer the user's questions.",
		OutputSchema: outputSchema,
	})

	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(routerAgent),
	}

	l := full.NewLauncher()

	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
