package git

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Commit struct {
	Hash    string
	Author  string
	Email   string
	Message string
	Date    time.Time
}

var ErrStop = errors.New("stop iteration")

func CollectCommits(repoPath string, limit int) ([]Commit, error) {
	fmt.Println("\nIniciando coleta de commits...")
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	commitIter, err := repo.Log(&git.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		return nil, err
	}

	var commits []Commit
	count := 0

	err = commitIter.ForEach(func(c *object.Commit) error {

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

	if err != nil && err != ErrStop {
		return nil, err
	}

	fmt.Println("Encontrados: ", len(commits), "commits\n")
	return commits, nil
}