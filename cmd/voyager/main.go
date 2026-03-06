package main

import (
	"context"
	"fmt"
	"strings"
	"voyager/internal/git"
	"voyager/internal/llm"
)

func main() {

	// LEITURA DE COMMITS
	commits, err := git.CollectCommits(".", 10)
	if err != nil {
		panic(err)
	}

	var builder strings.Builder

	for _, c := range commits {
		fmt.Fprintf(&builder, "Hash: %s\nMessage: %s---\n", c.Hash, c.Message)
	}

	commitsString := builder.String()

	// CARREGAMENTO DO PROMPT CLASSIFICADOR DE TIPOS
	prompt, err := llm.LoadPrompt("commit_analyzer_type.md")

	if err != nil {
		panic(err)
	}

	prompt = strings.ReplaceAll(prompt, "{{commits}}", commitsString)

	// CÓDIGO DO PROMPT

	client := llm.NewClient()

	err = client.Start()
	if err != nil {
		panic(err)
	}

	defer client.Stop()

	resp, err := client.Prompt(
		context.Background(),
		"mistral",
		prompt,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)


}