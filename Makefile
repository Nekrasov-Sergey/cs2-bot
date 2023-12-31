PWD = ${CURDIR}
# Название сервиса
SERVICE_NAME = cs2_bot
# Директория с proto файлами
PROTO_FILES_DIR = api
# Директория для pb файлов
GO_OUT_DIR = pkg/$(SERVICE_NAME)-api
# Proto файлы на компиляцию
PROTO_FILES = $(SERVICE_NAME)-api.proto
# 8 символов последнего коммита
LAST_COMMIT_HASH = $(shell git rev-parse HEAD | cut -c -8)

.SILENT:

.DEFAULT_GOAL := run

# example usage: `make sql/add_email_to_users`
sql/%:
	goose -dir migrations create $* sql

# Запуск сервиса
.PHONY: build
build:
	docker-compose build

# Запуск сервиса
.PHONY: run
run:
	docker-compose up -d

# Запустить тесты
.PHONY: test
test:
	go test $(PWD)/... -coverprofile=cover.out

# Запустить миграции
.PHONY: migrate
migrate:
	go run $(PWD)/cmd/migrate


# Компиляция proto файлов
.PHONY: generate
generate:
	mkdir -p $(PWD)/pkg/$(SERVICE_NAME)-api && \
	cd $(PROTO_FILES_DIR) && \
	protoc -I. --go_out=$(PWD)/$(GO_OUT_DIR) --go-grpc_out=$(PWD)/$(GO_OUT_DIR) $(PROTO_FILES) && \
	echo "New pb files generated"
