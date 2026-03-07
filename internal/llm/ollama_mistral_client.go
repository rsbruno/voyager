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
		http: &http.Client{}, // sem timeout aqui
		model: "mistral",
	}
}

func (c *Client) Start() error {
	if c == nil {
		return fmt.Errorf("llm: client é nil")
	}

	fmt.Println("Iniciando Ollama...")

	c.cmd = exec.Command("ollama", "serve")

	if err := c.cmd.Start(); err != nil {
		return fmt.Errorf("llm: erro ao iniciar ollama serve: %w", err)
	}

	if err := c.waitForServer(); err != nil {
		return err
	}

	fmt.Println("Ollama pronto")

	return nil
}

func (c *Client) waitForServer() error {
	for i := 0; i < 30; i++ {

		resp, err := c.http.Get("http://localhost:11434/api/tags")
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("llm: ollama não iniciou dentro do tempo esperado")
}

func (c *Client) Execute(prompt string) (string, error) {

	if c == nil {
		return "", fmt.Errorf("llm: client é nil")
	}

	if prompt == "" {
		return "", fmt.Errorf("llm: prompt vazio")
	}

	fmt.Println("Executando o modelo", c.model, "(aguarde)...")

	body := request{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("llm: erro ao serializar request: %w", err)
	}

	// timeout grande para permitir execução lenta da LLM local
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", fmt.Errorf("llm: erro ao criar request HTTP: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("llm: erro ao executar request HTTP: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("llm: resposta inválida da API (%d)", resp.StatusCode)
	}

	var result response

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("llm: erro ao decodificar resposta da API: %w", err)
	}

	if result.Response == "" {
		return "", fmt.Errorf("llm: resposta vazia do modelo")
	}

	return result.Response, nil
}

func (c *Client) Stop() error {

	fmt.Println("Encerrando Ollama...")

	if err := exec.Command("ollama", "stop", c.model).Run(); err != nil {
		return fmt.Errorf("llm: erro ao parar modelo %s: %w", c.model, err)
	}

	if err := exec.Command("pkill", "-f", "ollama").Run(); err != nil {
		return fmt.Errorf("llm: erro ao finalizar processo ollama: %w", err)
	}

	fmt.Println("Ollama encerrado")

	return nil
}