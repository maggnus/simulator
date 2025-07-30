
.PHONY: db

init:
	go run -mod=mod entgo.io/ent/cmd/ent new Record

db: ## Generate DB from schema
	go get entgo.io/ent/cmd/internal/printer
	go get entgo.io/ent/cmd/ent
	@find ./ent/* -maxdepth 0 -not -path "./ent/schema" -exec rm -r {} +
	go run entgo.io/ent/cmd/ent generate --feature sql/modifier --feature sql/upsert --feature sql/execquery ./ent/schema/
	#go generate ./ent

doc: ## Generate documentation
	go install golang.org/x/tools/cmd/godoc@latest
	godoc -http=:8080

TAG ?= latest

build: ## Build
	docker-compose build --build-arg TAG=${TAG}

up: ## Up
	docker-compose up -d  --remove-orphans

down: ## Down
	docker-compose down

.DEFAULT_GOAL := db
