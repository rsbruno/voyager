package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os/exec"
	"fmt"
)

type Client struct {
	baseURL string
	http    *http.Client
	cmd     *exec.Cmd
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
		http:    &http.Client{},
	}
}

func (c *Client) Start() error {
	c.cmd = exec.Command("ollama", "serve")
	fmt.Println("Iniciado Ollama")
	
	return c.cmd.Start()
}


func (c *Client) Execute(prompt string) (string, error) {
	model := "mistral"

	fmt.Println("Executando o modelo", model, "(Aguarde a finalização)...")

	body := request{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	data, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		c.baseURL,
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result response

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Response, nil
}

func (c *Client) Stop() {
	fmt.Println("Encerrando Ollama...")

	exec.Command("ollama", "stop", "mistral").Run()

	exec.Command("pkill", "-f", "ollama").Run()
}