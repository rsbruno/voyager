package main

import (
	"context"
	"fmt"
	"voyager/internal/git"
	"voyager/internal/llm"
	"voyager/internal/analyser"
)

func main() {

	// LEITURA DE COMMITS
	commits, err := git.CollectCommits(".", 10)
	if err != nil {
		panic(err)
	}

	client := llm.NewClient()

	err = client.Start()
	if err != nil {
		panic(err)
	}

	defer client.Stop()

	// CLASSIFICAÇÃO DE TIPOS DE COMMITS
	resp, err := analyser.CommitAnalyzerType(context.Background(), commits, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}