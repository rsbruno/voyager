package infra

import (
	"context"
	"fmt"
	"os"
	"sync"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var (
	sheetsService *sheets.Service
	once          sync.Once
	initErr       error
)

func GetSheetsService() (*sheets.Service, error) {
	once.Do(func() {

		ctx := context.Background()

		credBytes, err := os.ReadFile("credentials.json")
		if err != nil {
			initErr = fmt.Errorf("Erro ao ler credentials.json: %w", err)
			return
		}

		config, err := google.JWTConfigFromJSON(
			credBytes,
			sheets.SpreadsheetsScope,
		)
		if err != nil {
			initErr = fmt.Errorf("Erro ao criar config JWT: %w", err)
			return
		}

		client := config.Client(ctx)

		sheetsService, err = sheets.NewService(
			ctx,
			option.WithHTTPClient(client),
		)
		if err != nil {
			initErr = fmt.Errorf("Erro ao inicializar serviço Google Sheets: %w", err)
			return
		}
	})

	if initErr != nil {
		return nil, initErr
	}

	return sheetsService, nil
}

func Write(sheetID string, cell string, value string) error {

	service, err := GetSheetsService()
	if err != nil {
		return fmt.Errorf("Erro ao obter serviço sheets: %w", err)
	}

	values := [][]interface{}{
		{value},
	}

	vr := &sheets.ValueRange{
		Values: values,
	}

	_, err = service.Spreadsheets.Values.Update(
		sheetID,
		cell,
		vr,
	).ValueInputOption("RAW").Do()

	if err != nil {
		return fmt.Errorf(
			"Erro ao escrever na planilha (sheetID=%s cell=%s): %w",
			sheetID,
			cell,
			err,
		)
	}

	return nil
}

func SheetExists(sheetID, sheetName string) (bool, error) {

	service, err := GetSheetsService()
	if err != nil {
		return false, fmt.Errorf("Erro ao obter serviço sheets: %w", err)
	}

	resp, err := service.Spreadsheets.Get(sheetID).Do()
	if err != nil {
		return false, fmt.Errorf(
			"Erro ao consultar planilha (sheetID=%s): %w",
			sheetID,
			err,
		)
	}

	for _, s := range resp.Sheets {
		if s.Properties.Title == sheetName {
			return true, nil
		}
	}

	return false, nil
}