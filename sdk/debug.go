package sdk

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Debug[T any](filename string) (*T, error) {
	fmt.Printf("Método em modo debug, retornado o mock: %s\n", filename)

	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não encontrado")
	}

	if os.Getenv("MODE") != "development" {
		return nil, nil
	}

	path := filepath.Join("sdk","data", filename)

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var result T

	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}