# Voyager

Ferramenta em Go que coleta os commits de um repositório git, usa o Gemini para
classificar o tipo de cada commit e padronizar sua mensagem, e registra o
resultado em uma planilha do Google Sheets.

## Como funciona

1. Coleta commits de um repositório local, filtrando por autor e período.
2. Envia os commits para o Gemini para:
   - classificar o tipo (feat, fix, refactor, ...);
   - reescrever a mensagem em um formato padronizado.
3. Agrupa os commits por data e escreve o resultado em uma planilha do Google
   Sheets (uma aba por mês, no formato `MM/YYYY`, já existente na planilha).

## Requisitos

- Go 1.25+
- Uma [service account do Google Cloud](https://cloud.google.com/iam/docs/service-account-overview)
  com acesso à API do Google Sheets, salva como `credentials.json` na raiz do
  projeto, com a planilha compartilhada com o e-mail dessa service account.
- Credenciais do Gemini configuradas via
  [Application Default Credentials](https://ai.google.dev/gemini-api/docs)
  ou variável de ambiente esperada pelo SDK `google.golang.org/genai`.

## Configuração

Copie `.env.example` para `.env` e preencha as variáveis:

```
cp .env.example .env
```

| Variável          | Obrigatória | Descrição                                               |
| ----------------- | ----------- | --------------------------------------------------------- |
| `REPO_PATH`       | sim         | Caminho absoluto do repositório git a ser analisado       |
| `AUTHOR_EMAIL`    | sim         | E-mail do autor cujos commits serão coletados             |
| `SINCE_DATE`      | sim         | Data inicial da coleta (`YYYY-MM-DD`)                      |
| `GOOGLE_SHEET_ID` | sim         | ID da planilha do Google Sheets                            |
| `GEMINI_MODEL`    | não         | Modelo Gemini a usar (padrão: `gemini-3-flash-preview`)    |
| `MODE`            | não         | `development` usa os mocks de `sdk/data/` no lugar da LLM  |

`credentials.json` e `.env` nunca devem ser commitados (já estão no
`.gitignore`).

## Executando

```
go run ./cmd/voyager
```
