# games-api

Exemplo de API em Go organizanda por *feature folders*, usando Fiber. Dados de seed (80 jogos) são lidos de `data/games_seed.json`.


## Requisitos

- Go 1.21+
- Docker & Docker Compose (se for usar containers)
- make (opcional, recomendado)

## Estrutura do projeto

```
games-api
│   .dockerignore
│   .gitignore
│   docker-compose.yml
│   Dockerfile
│   go.mod
│   go.sum
│   main.go
│   Makefile
│   README.md
│
├───data
│       games_seed.json
│
├───internal
│   └───games
│           handler.go
│           model.go
│           repository.go
│           routes.go
│           service.go
│
└───pkg
    └───ulidutil
            ulid.go
```

## Como rodar (local, sem Docker)

1. Instale dependências e verifique módulos:
```bash
make tidy
```

2. Rodar localmente (usa `data/games_seed.json`):
```bash
make run
# ou
go run main.go
```

3. Build local do binário:
```bash
make build
# binário gerado em bin/games-api
```

## Rodando com Docker / Docker Compose
1. Construir a imagem Docker:
```bash
make docker-build
```

2. Subir containers (monta ./data como /app/data no container):
```bash
make up
```

3. Logs, parar e remover:
```bash
make logs
make down
```

> **Observação:** o `docker-compose.yml` monta a pasta local `./data` em `/app/data` do container (read-only). Garanta que `data/games_seed.json` esteja presente na raiz do projeto antes de subir o compose.

## Makefile — comandos úteis

Lista dos comandos principais (já disponíveis no `Makefile`):

- `make` ou `make help` — mostra os comandos disponíveis.
- `make tidy` — `go mod tidy` e `go mod verify`.
- `make build` — compila o binário em `bin/games-api`.
- `make run` — executa `go run main.go` (usa o seed JSON em ./data).
- `make test` — roda `go test ./... -v`.
- `make lint` — formata o código (`go fmt ./...`).

**Docker related:**
- `make docker-build` — builda a imagem Docker localmente.
- `make up` — `docker-compose up --build -d`.
- `make down` — `docker-compose down`.
- `make logs` — segue logs do serviço.
- `make shell` — abre um shell dentro do container (`/bin/sh`).
- `make docker-clean` — remove containers/imagens criadas.

**Utilitários:**
- `make health` — `curl` no endpoint `GET /games` (usa `PORT` 3000 por padrão).

## Seed data

- O seed com 80 jogos está em `data/games_seed.json` e é versionado no repositório.
- O serviço carrega esse arquivo ao iniciar; a aplicação espera encontrar `data/games_seed.json` no diretório de trabalho.

## Exemplo de requisição CURL (local)

Você pode testar a API localmente (porta padrão 3000) com todos os parâmetros mockados:

```bash
curl -X GET "http://localhost:3000/games?name=mario&platform=snes&gender=Platform&subGender=Adventure" \
-H "Accept: application/json" | jq
```

Exemplo de resposta:

```json
{
    "data": [
        {
            "id": "01H1Z4Q0V1F2G3H4J5K6L7M8N",
            "name": "Super Mario World",
            "releaseDate": "1990-11-21",
            "platform": "snes",
            "gender": "Platform",
            "subGender": "Adventure",
            "rating": 9.0
        }
    ]
}
```


Parâmetros de exemplo usados:
- `name=mario` → busca por nome contendo "mario" (case-insensitive)
- `platform=snes` → filtra apenas jogos de Super Nintendo
- `gender=Platform` → gênero principal
- `subGender=Adventure` → subgênero


Todos os resultados vêm **ordenados por rating (desc)** por padrão.