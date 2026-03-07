Commit Type Classifier

Classify Git commit messages following the Conventional Commits specification and map each type to its corresponding short Portuguese label.

Types and Labels

feat: Criado
fix: Ajustado
bug: Ajustado
config: Configurado
refactor: Refatorado
style: Estilizado
docs: Documentado
test: Testado
perf: Otimizado
chore: Mantido
build: Construído
ci: Integrado
revert: Revertido

Rules

Identify the commit type from the message.

Each commit will be provided in a list.

Commits are separated by ---.

Return only JSON.

Do not include explanations.

Do not include extra text.

Output Format

The response must always be a JSON object.

Each key must be the commit hash.

Each value must be the corresponding Portuguese label.

Property names must be in English.

Values must be in Brazilian Portuguese (pt-BR).

Do not translate technical terms, programming keywords, library names, framework names, API names, commit hashes, code identifiers, or English terminology commonly used in software development. Preserve them exactly as written (e.g., commit, refactor, API, client, server, JSON, HTTP, Git, Docker, Kubernetes, Go, function, struct, interface).

Preserve code identifiers, variable names, and function names exactly as written.

Example:

{
  "80fc51611ffc8610010e58aca29ad5b57a1b24b3": "Criado",
  "7bdae42fdea507c76352ef596b06b0e25c93b3b9": "Ajustado"
}

Commits:
{{commits}}