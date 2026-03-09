package main

import (
	"fmt"
	"log"
	"time"
	"voyager/internal/analyser"
	"voyager/internal/git"
	"voyager/internal/llm"
	"voyager/internal/reports"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não encontrado")
	}

	repo := "/home/bruno-santos/Documentos/trocafone/trocafone-new-trade-in-plataform"

	since := time.Date(2026, 3, 1, 0, 0, 0, 0, time.Local)
	until := time.Now()

	gitCommits, err := git.CollectCommits(repo, "rsbruno.cdc@gmail.com", since, until)
	if err != nil {
		panic(err)
	}

	client, err := llm.NewGeminiClient()
	if err != nil {
		panic(err)
	}

	commits := llm.PromptCommitBuilder(gitCommits)

	types, err := analyser.CommitAnalyzerType(commits, client)
	if err != nil {
		panic(err)
	}

	messages, err := analyser.CommitAnalyzerMessage(commits, client)
	if err != nil {
		panic(err)
	}

	workCommits := make([]git.Commit, len(gitCommits))

	for i, c := range gitCommits {

		hash := c.Hash

		workCommits[i] = git.Commit{
			Hash:    hash,
			Author:  c.Author,
			Email:   c.Email,
			Date:    c.Date,
			Message: messages[hash],
			Type: types[hash],
		}

	}

	err = reports.ReportCommitsVoyagerSheets(workCommits)
	if err != nil {
		fmt.Println(err)
	}
}