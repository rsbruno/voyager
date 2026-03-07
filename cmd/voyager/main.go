package main

import (
	"fmt"
	"voyager/internal/analyser"
	"voyager/internal/git"
	"voyager/internal/llm"
	"voyager/internal/reports"
)


func main() {
	gitCommits, err := git.CollectCommits(".", 10)
	if err != nil {
		panic(err)
	}

	client := llm.NewClient()
	err = client.Start()
	if err != nil {
		panic(err)
	}
	defer client.Stop()


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
			Type:    types[hash],
			Message: messages[hash],
		}
	}

	err = reports.ReportCommitsVoyagerSheets(workCommits)
	if err != nil {
		fmt.Println(err)
	}
}