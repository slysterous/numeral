
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

lint: ## perform linting
	golint -set_exit_status=1 `go list ./...`
	go vet -mod vendor -tags=integration ./...

fmt: ## gofmt all files excluding vendor
	@go fmt ./...

test: ## run tests
	go test ./... -race -cover -v -count 1 -p 1 -tags=integration -mod vendor

ci: ## run ci
	./scripts/ci.sh