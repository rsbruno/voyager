Commit Message Analyzer

Analyze Git commit messages and rewrite them as standardized messages following these rules:

Rules

- Each message must have exactly 20 words.
- Messages must be written in the past tense.
- Messages must **not start with "Eu"**; remove the pronoun if present.
- Remove any leading spaces and capitalize the first letter of the sentence.
- Do not include commit type labels (like 'Adicionou:', 'Refatorou:'); use only the verb phrase.
- Messages must not be generic; keep context precise.
- Do not hallucinate; preserve the original commit context.
- Do not detail file-level changes; focus only on general scope.
- Messages must be written in Brazilian Portuguese (pt-BR).
- Do not translate technical terms, programming keywords, library names, framework names, API names, commit hashes, code identifiers, or English terminology commonly used in software development. Preserve them exactly as written (e.g., commit, refactor, API, client, server, JSON, HTTP, Git, Docker, Kubernetes, Go, function, struct, interface).
- Preserve code identifiers, variable names, and function names exactly as written.

Return only JSON.

Do not include explanations.

Do not include extra text.

Output Format

The response must always be a JSON object.

Each key must be the commit hash.

Each value must be the rewritten message following the rules above.

Property names must be in English.

Example:

{
  "7fb507a8293fa4896dd0d1887d017aa15d6a73d4": "Construi o construtor de prompts e simplifiquei a interface do cliente LLM ao centralizar a carga de prompts e commits.",
  "75a6b1a59935411e50bb9c58ff4b2ee2c8fcd3d2": "Implementei análise de tipo de commit baseada em LLM e refatorei o principal para usar o novo módulo de análise."
}

Commits:
{{commits}}