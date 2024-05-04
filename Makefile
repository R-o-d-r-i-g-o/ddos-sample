.PHONY: clean-cache

install:
	@echo Validating dependecies...
	@go mod tidy
	@echo Generating vendor...
	@go mod vendor

clean-cache:
	@echo Cleaning golang cache...
	@go clean -modcache
	@echo Deleting vendor folder if exists...
	@if exists .\vendor rmdir /s /q .\vendor

format:
	@gofmt -w .

run:
	@go run ./cmd/main.go