package main

import (
	"fmt"
	"voyager/internal/git"
	"voyager/internal/reports"
)


func main() {
	commits, err := git.CollectCommits(".", 10)
	if err != nil {
		panic(err)
	}

	//promptCommits := llm.PromptCommitBuilder(commits)

	types := map[string]string{
		"351802a1a4bfdf7498fc3858c7d4221068e378fa":"Criado",
		"75a6b1a59935411e50bb9c58ff4b2ee2c8fcd3d2":"Refatorado",
		"7fb507a8293fa4896dd0d1887d017aa15d6a73d4":"Refatorado",
		"80fc51611ffc8610010e58aca29ad5b57a1b24b3":"Criado",
		"7bdae42fdea507c76352ef596b06b0e25c93b3b9": "Criado",
		"babea92cf7f3fa0a3cf2fb85cd39e59d582179a3":"Refatorado",
		"ddd77c7126f0d5ce38f385f579c896d58d4465f1":"Refatorado",
	}

	messages := map[string]string{
		"351802a1a4bfdf7498fc3858c7d4221068e378fa": "Adiciona utilitário de coleta de commits Git e exibe commits recentes de um repositório.",
		"75a6b1a59935411e50bb9c58ff4b2ee2c8fcd3d2": "Refatorou principal para usar novo módulo de análise de tipos de commit baseado em LLM.",
		"7bdae42fdea507c76352ef596b06b0e25c93b3b9": "Inicia projeto Voyager com módulo Go, comando principal e pacotes internos para Git, LLM e análise de commits.",
		"7fb507a8293fa4896dd0d1887d017aa15d6a73d4": "Refatorou interface do cliente LLM centralizando carga de prompts e formatagem de commits.",
		"80fc51611ffc8610010e58aca29ad5b57a1b24b3": "Inicia projeto Voyager com módulo Go, comando principal e pacotes internos para Git, LLM e análise de commits.",
		"babea92cf7f3fa0a3cf2fb85cd39e59d582179a3": "Refatorou arquivo de prompt de análise de tipo de commit e atualizou caminho de carregamento.",
		"ddd77c7126f0d5ce38f385f579c896d58d4465f1": "Construi construidor de mensagens de commit usando LLM e integre-o em main.go.",
	}

	workCommits := make([]git.Commit, len(commits))
	for i, c := range commits {
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
	



	// client := llm.NewClient()
	// err = client.Start()
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Stop()

	// types, err := analyser.CommitAnalyzerType(promptCommits, client)
	// if err != nil {
	// 	panic(err)
	// }

	// messages, err := analyser.CommitAnalyzerMessage(promptCommits, client)
	// if err != nil {
	// 	panic(err)
	// }

	//finalCommits := make([]git.Commit, len(commits))

	// for i, c := range commits {
	// 	hash := c.Hash

	// 	commitType := types[hash]

	// 	if commitType == "" {
	// 		commitType = ""
	// 	}


	// 	finalCommits[i] = git.Commit{
	// 		Hash:    hash,
	// 		Author:  c.Author,
	// 		Email:   c.Email,
	// 		Date:    c.Date,
	// 		Type:    commitType,
	// 	}
	// }




// 	// 5. Obtem mensagens de commit do LLM
// 	respMessages, err := analyser.CommitAnalyzerMessage(promptCommits, client)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 6. Decodifica JSON da LLM em map[string]string
// 	var typesJSON map[string]string
// 	if err := json.Unmarshal([]byte(respTypes), &typesJSON); err != nil {
// 		panic(fmt.Errorf("falha ao decodificar typesJSON: %w", err))
// 	}

// 	var messagesJSON map[string]string
// 	if err := json.Unmarshal([]byte(respMessages), &messagesJSON); err != nil {
// 		panic(fmt.Errorf("falha ao decodificar messagesJSON: %w", err))
// 	}

// 	// 7. Monta o slice final de commits
// 	finalCommits := make([]git.Commit, len(commits))
// 	for i, c := range commits {
// 		hash := c.Hash

// 		commitType := typesJSON[hash]
// 		msg := messagesJSON[hash]

// 		// fallback caso LLM não tenha retornado algo
// 		if commitType == "" {
// 			commitType = ""
// 		}
// 		if msg == "" {
// 			msg = c.Message
// 		}


// 		finalCommits[i] = git.Commit{
// 			Hash:    hash,
// 			Author:  c.Author,
// 			Email:   c.Email,
// 			Date:    c.Date,
// 			Type:    commitType,
// 		}
// 	}

	// 8. Exibe o resultado final

}