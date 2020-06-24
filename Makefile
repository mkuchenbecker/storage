.PHONY: lint
lint: 
	golangci-lint run

.PHONY: install-golang-ci
lint-ci:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh

.PHONY: generate
generate:
	protoc --proto_path=./api/ --go_out=plugins=grpc:./api/ --go_opt=paths=source_relative wire.proto  

.PHONY: integration
integration:
	richgo test -race -cover -tags=integration ./...
	
.PHONY: build
build:
	docker build -t mkuchenbecker/storage:storage-latest -f service.dockerfile . \

.PHONY: push
push:
	docker push mkuchenbecker/storage:storage-latest

.PHONY: fmt
fmt:
	@echo "fmt:"
	scripts/fmt

