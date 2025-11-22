package main

import (
	"fmt"

	"google.golang.org/adk/tool"
)

type getWeatherArgs struct {
	City string `json:"city" jsonschema:"The city to get weather for."`
}

func getWeather(ctx tool.Context, args getWeatherArgs) (string, error) {
	fmt.Printf("[Tool] Getting weather for %s...\n", args.City)
	return fmt.Sprintf("The weather in %s is Sunny, 25°C", args.City), nil
}
