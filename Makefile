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

PHONY: sam, sam-run, default

