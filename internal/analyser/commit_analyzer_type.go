package analyser

import (
	"context"
	"fmt"
	"strings"

	"voyager/internal/llm"
	"voyager/internal/git"
)

// CommitAnalyzerType envia os commits para a LLM e retorna a resposta
func CommitAnalyzerType(ctx context.Context, commits []git.Commit, client *llm.Client) (string, error) {
	fmt.Println("\nIniciado a classificação de tipos dos commits...")
	var builder strings.Builder

	for _, c := range commits {
		fmt.Fprintf(&builder, "Hash: %s\nMessage: %s---\n", c.Hash, c.Message)
	}

	commitsString := builder.String()

	// CARREGAMENTO DO PROMPT CLASSIFICADOR DE TIPOS
	prompt, err := llm.LoadPrompt("commit_analyzer_type.md")
	if err != nil {
		return "", err
	}

	prompt = strings.ReplaceAll(prompt, "{{commits}}", commitsString)

	// CÓDIGO DO PROMPT usando o client recebido
	resp, err := client.Prompt(ctx, "mistral", prompt)
	if err != nil {
		return "", err
	}

	fmt.Println("Classificação de tipos dos commits finalizada!\n")
	// Retorna a resposta bruta da LLM
	return resp, nil
}