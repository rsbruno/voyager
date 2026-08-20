package analyser

import (
	"fmt"
	"strings"

	"voyager/internal/llm"
	"voyager/internal/utils"
	"voyager/sdk"
)

func CommitAnalyzerMessage(commits string, client *llm.GeminiClient) (map[string]string, error) {
	fmt.Println("\nIniciado a classificação de mensagens dos commits...")

	data, err := sdk.Debug[map[string]string]("messages.json")
	if data != nil {
		return *data, err
	}

	if client == nil {
		return nil, fmt.Errorf("Client LLM é nil")
	}

	if commits == "" {
		return nil, fmt.Errorf("Lista de commits vazia")
	}

	prompt, err := llm.LoadPrompt("commit_analyzer_message.md")
	if err != nil {
		return nil, fmt.Errorf(
			"Erro ao carregar prompt commit_analyzer_message.md: %w",
			err,
		)
	}

	prompt = strings.ReplaceAll(prompt, "{{commits}}", commits)

	resp, err := client.Execute(prompt)
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar LLM: %w", err)
	}

	parsed, err := utils.JsonParse(resp)
	if err != nil {
		fmt.Println(resp)
		return nil, fmt.Errorf(
			"Erro ao fazer parse do JSON retornado pela LLM (message): %w",
			err,
		)
	}

	fmt.Println("Classificação de mensagens dos commits finalizada!\n")

	return parsed, nil
}