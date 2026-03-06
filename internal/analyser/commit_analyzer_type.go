package analyser

import (
	"fmt"
	"strings"

	"voyager/internal/llm"
)

func CommitAnalyzerType(commits string, client *llm.Client) (string, error) {
	fmt.Println("\nIniciado a classificação de tipos dos commits...")

	prompt, err := llm.LoadPrompt("commit_analyzer_type.md")
	if err != nil {
		return "", err
	}

	prompt = strings.ReplaceAll(prompt, "{{commits}}", commits)

	resp, err := client.Execute(prompt)
	if err != nil {
		return "", err
	}

	fmt.Println("Classificação de tipos dos commits finalizada!\n")
	return resp, nil
}