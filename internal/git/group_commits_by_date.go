package git

func GroupCommitsByDate(commits []Commit) map[string][]Commit {
	grouped := make(map[string][]Commit)

	for _, c := range commits {
		dateKey := c.Date.Format("2006-01-02")

		grouped[dateKey] = append(grouped[dateKey], c)
	}

	return grouped
}