package utils

import (
	"encoding/json"
	"fmt"
)

func JsonParse(jsonStr string) (map[string]string, error) {
	var parsed map[string]string
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		return nil, fmt.Errorf("Erro ao decodificar json: %w", err)
	}
	return parsed, nil
}