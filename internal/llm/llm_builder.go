package llm

import (
	"path/filepath"
	"regexp"
	"fmt"
	"os"
	"strings"
	"time"

	"voyager/internal/git"
)


var importRegex = regexp.MustCompile(`{{import:(.*?)}}`)

func LoadPrompt(name string) (string, error) {
	fmt.Println("Carregando prompt: ", name)

	path := filepath.Join("internal", "prompts", name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("prompt não encontrado: %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("erro ao ler o prompt: %w", err)
	}

	fmt.Println("Prompt carregado com sucesso!")
	return string(data), nil
}

func PromptCommitBuilder(commits []git.Commit) string {
	var builder strings.Builder

	for _, c := range commits {
		fmt.Fprintf(
			&builder,
			"Hash: %s\nAuthor: %s\nEmail: %s\nDate: %s\nMessage: %s\n---\n",
			c.Hash,
			c.Author,
			c.Email,
			c.Date.Format(time.RFC3339),
			c.Message,
		)
	}

	return builder.String()
}