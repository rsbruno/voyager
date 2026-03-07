package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

type Client struct {
	baseURL string
	http    *http.Client
	cmd     *exec.Cmd
	model   string
}

type request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type response struct {
	Response string `json:"response"`
}

func NewClient() *Client {
	return &Client{
		baseURL: "http://localhost:11434/api/generate",
		http: &http.Client{
			Timeout: 120 * time.Second,
		},
		model: "mistral",
	}
}

func (c *Client) Start() error {
	if c == nil {
		return fmt.Errorf("Client é nil")
	}

	c.cmd = exec.Command("ollama", "serve")

	fmt.Println("Iniciando Ollama...")

	if err := c.cmd.Start(); err != nil {
		return fmt.Errorf("Erro ao iniciar ollama serve: %w", err)
	}

	fmt.Println("Ollama iniciado")

	return nil
}

func (c *Client) Execute(prompt string) (string, error) {
	if c == nil {
		return "", fmt.Errorf("Client é nil")
	}

	if prompt == "" {
		return "", fmt.Errorf("Prompt vazio")
	}

	fmt.Println("Executando o modelo", c.model, "(aguarde)...")

	body := request{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("Erro ao serializar request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		c.baseURL,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", fmt.Errorf("Erro ao criar request HTTP: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("Erro ao executar request HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Resposta inválida da API (%d)", resp.StatusCode)
	}

	var result response

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("Erro ao decodificar resposta da API: %w", err)
	}

	if result.Response == "" {
		return "", fmt.Errorf("Resposta vazia do modelo")
	}

	return result.Response, nil
}

func (c *Client) Stop() error {
	fmt.Println("Encerrando Ollama...")

	if err := exec.Command("ollama", "stop", c.model).Run(); err != nil {
		return fmt.Errorf("Erro ao parar modelo %s: %w", c.model, err)
	}

	if err := exec.Command("pkill", "-f", "ollama").Run(); err != nil {
		return fmt.Errorf("Erro ao finalizar processo ollama: %w", err)
	}

	fmt.Println("Ollama encerrado")

	return nil
}