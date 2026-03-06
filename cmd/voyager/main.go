package main

import (
	"fmt"

	"voyager/internal/git"
)

func main() {

	commits, err := git.CollectCommits(".", 10)
	if err != nil {
		panic(err)
	}

	for _, c := range commits {
		fmt.Println("----")
		fmt.Println("Hash:", c.Hash)
		fmt.Println("Author:", c.Author)
		fmt.Println("Email:", c.Email)
		fmt.Println("Message:", c.Message)
		fmt.Println("Date:", c.Date)
	}
}