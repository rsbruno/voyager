package reports

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"voyager/internal/git"
	"voyager/internal/infra"
	"voyager/internal/utils"
)


func ReportCommitsVoyagerSheets(commits []git.Commit) error {
	fmt.Println("\nRegistrando os commits no Google Sheet (Voyager Sheets)")

	grouped := git.GroupCommitsByDate(commits)
	sheetId := "1w0jHDDU5onnzQ1RhoQoq4wkg18LsXGs8vnN_N-G68mg"

	var errs []error

	for date, commits := range grouped {

		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			errs = append(errs, fmt.Errorf("Erro ao converter data %s: %w", date, err))
			continue
		}

		sheetName := t.Format("01/2006")
		row := t.Day() + 1

		exists, err := infra.SheetExists(sheetId, sheetName)
		if err != nil {
			errs = append(errs, fmt.Errorf("Erro verificando aba %s: %w", sheetName, err))
			continue
		}

		if !exists {
			errs = append(errs, fmt.Errorf("Aba %s não existe na planilha", sheetName))
			continue
		}

		var activity []string
		var subject []string

		for _, c := range commits {
			activity = append(activity, c.Message)
			subject = append(subject, fmt.Sprintf("[Tarefa][Nº][%s][%s]", utils.ShortHash(c.Hash), c.Type))
		}

		err = infra.Write(
			sheetId,
			fmt.Sprintf("%s!N%d", sheetName, row),
			strings.Join(subject, "\n"),
		)
		if err != nil {
			errs = append(errs, fmt.Errorf("Erro escrevendo assunto %s linha %d: %w", sheetName, row, err))
		}

		err = infra.Write(
			sheetId,
			fmt.Sprintf("%s!O%d", sheetName, row),
			strings.Join(activity, "\n"),
		)
		if err != nil {
			errs = append(errs, fmt.Errorf("Erro escrevendo atividade %s linha %d: %w", sheetName, row, err))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	fmt.Println("Google Sheet (Voyager Sheets) atualizado com sucesso")
	return nil
}