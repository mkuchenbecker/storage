.PHONY: lint
lint: 
	golangci-lint run

.PHONY: install-golang-ci
lint-ci:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh

.PHONY: generate
generate:
	protoc --proto_path=api/ \
	--include_imports \
    --include_source_info \
	--go_out=plugins=grpc:./api/ \
	--go_opt=paths=source_relative \
    --descriptor_set_out=./api/wire_descriptor.pb \
	wire.proto  
	protoc --proto_path=./service/datamodel \
	--go_out=plugins=grpc:./service/datamodel \
	--go_opt=paths=source_relative \
	datamodel.proto
	protoc --proto_path=./testing/model \
	--go_out=plugins=grpc:./testing/model \
	--go_opt=paths=source_relative \
	foo.proto
	protoc --proto_path=universe/vector/ \
	--go_out=plugins=grpc:./universe/vector/ \
	--go_opt=paths=source_relative \
	vector.proto  

.PHONY: init
init: generate build
	echo "Initialized"

.PHONY: integration
integration:
	richgo test -race -cover -tags=inFtegration ./...

.PHONY: unit
unit:
	richgo test -race -cover ./...
	
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

