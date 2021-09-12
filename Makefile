
default:
	@go test -v ./...
	@go build ./...

sam:
	@go test -v ./...
	@sam build --template ./infra/aws/sam/template.yaml

sam-run:
	@go test -v ./...
	@sam build --template ./infra/aws/sam/template.yaml
	@sam local start-api

PHONY: sam, sam-run, default

