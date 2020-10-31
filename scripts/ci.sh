#!/bin/sh
set -e

./scripts/gofmtcheck.sh

# lint
echo "==> Checking lint"
$GOPATH/bin/golint -set_exit_status=1 `go list ./...`

# vet
echo "==> Checking vet"
go vet ./...

# test
echo "==> Running all tests with race detector and coverage (this will take a while)"
go test ./... -race -p 1 -cover -v -tags=integration