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

func CollectCommits(
	repoPath string,
	authorEmail string,
	since time.Time,
	until time.Time,
) ([]Commit, error) {

	if repoPath == "" {
		return nil, fmt.Errorf("repoPath não pode ser vazio")
	}

	fmt.Println("\nIniciando coleta de commits...")

	repo, err := gogit.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir repositório (%s): %w", repoPath, err)
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("erro ao obter HEAD do repositório: %w", err)
	}

	commitIter, err := repo.Log(&gogit.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao iniciar log de commits: %w", err)
	}

	var commits []Commit

	err = commitIter.ForEach(func(c *object.Commit) error {

		if c == nil {
			return fmt.Errorf("commit nulo encontrado durante iteração")
		}

		if authorEmail != "" && c.Author.Email != authorEmail {
			return nil
		}

		commitDate := c.Author.When

		if !since.IsZero() && commitDate.Before(since) {
			return nil
		}

		if !until.IsZero() && commitDate.After(until) {
			return nil
		}

		commits = append(commits, Commit{
			Hash:    c.Hash.String(),
			Author:  c.Author.Name,
			Email:   c.Author.Email,
			Message: c.Message,
			Date:    commitDate,
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao iterar commits: %w", err)
	}

	fmt.Println("Encontrados:", len(commits), "commits")

	return commits, nil
}