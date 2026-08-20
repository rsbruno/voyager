package main

import (
	"fmt"
	"log"
	"os"
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

	repo := os.Getenv("REPO_PATH")
	if repo == "" {
		log.Fatal("REPO_PATH não definido (configure no .env)")
	}

	authorEmail := os.Getenv("AUTHOR_EMAIL")
	if authorEmail == "" {
		log.Fatal("AUTHOR_EMAIL não definido (configure no .env)")
	}

	sheetID := os.Getenv("GOOGLE_SHEET_ID")
	if sheetID == "" {
		log.Fatal("GOOGLE_SHEET_ID não definido (configure no .env)")
	}

	since, err := parseSince(os.Getenv("SINCE_DATE"))
	if err != nil {
		log.Fatal(err)
	}

	until := time.Now()

	gitCommits, err := git.CollectCommits(repo, authorEmail, since, until)
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
			Type:    types[hash],
		}

	}

	err = reports.ReportCommitsVoyagerSheets(workCommits, sheetID)
	if err != nil {
		fmt.Println(err)
	}
}

func parseSince(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, fmt.Errorf("SINCE_DATE não definido (configure no .env, formato YYYY-MM-DD)")
	}

	t, err := time.ParseInLocation("2006-01-02", value, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("SINCE_DATE inválido (%s): use o formato YYYY-MM-DD: %w", value, err)
	}

	return t, nil
}
