.PHONY: install build deploy

install:
	npm install -g aws-cdk
	go mod tidy

build:
	go build -o cdk-api-integrations

deploy:
	cdk deploy --all

run:
	go run ./cdk-api-integrations.go
