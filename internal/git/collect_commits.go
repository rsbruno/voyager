package git

import (
	"errors"
	"fmt"
	"time"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Commit struct {
	Hash    string
	Author  string
	Email   string
	Message string
	Date    time.Time
	Type    string
}

var ErrStop = errors.New("Stop commit iteration")

func CollectCommits(repoPath string, limit int) ([]Commit, error) {

	if repoPath == "" {
		return nil, fmt.Errorf("RepoPath não pode ser vazio")
	}

	fmt.Println("\nIniciando coleta de commits...")

	repo, err := gogit.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir repositório (%s): %w", repoPath, err)
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter HEAD do repositório: %w", err)
	}

	commitIter, err := repo.Log(&gogit.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		return nil, fmt.Errorf("Erro ao iniciar log de commits: %w", err)
	}

	var commits []Commit
	count := 0

	err = commitIter.ForEach(func(c *object.Commit) error {

		if c == nil {
			return fmt.Errorf("Commit nulo encontrado durante iteração")
		}

		commits = append(commits, Commit{
			Hash:    c.Hash.String(),
			Author:  c.Author.Name,
			Email:   c.Author.Email,
			Message: c.Message,
			Date:    c.Author.When,
		})

		count++

		if limit > 0 && count >= limit {
			return ErrStop
		}

		return nil
	})

	if err != nil && !errors.Is(err, ErrStop) {
		return nil, fmt.Errorf("Erro ao iterar commits: %w", err)
	}

	fmt.Println("Encontrados:", len(commits), "commits\n")

	return commits, nil
}