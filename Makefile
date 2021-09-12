SAM_TEMPLATE_PATH := ./aws-sam-cli/template.yaml

default:
	@go test -v ./...
	@go build ./...

sam:
	@go test -v ./...
	@sam build --template $(SAM_TEMPLATE_PATH)

sam-run:
	@go test -v ./...
	@sam build --template $(SAM_TEMPLATE_PATH)
	@sam local start-api

coverage:
	@go test -v -coverprofile cover.out ./...
	@go tool cover -html=cover.out -o cover.html
	@explorer.exe cover.html

PHONY: coverage, sam, sam-run, default

