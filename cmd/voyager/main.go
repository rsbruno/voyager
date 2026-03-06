package main

import (
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
	
	promptCommits := llm.PromptCommitBuilder(commits)

	client := llm.NewClient()

	err = client.Start()
	if err != nil {
		panic(err)
	}

	defer client.Stop()

	// resp, err := analyser.CommitAnalyzerType(promptCommits, client)
	// if err != nil {
	// 	panic(err)
	// }


	resp, err := analyser.CommitAnalyzerMessage(promptCommits, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}