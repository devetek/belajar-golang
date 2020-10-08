.PHONY: run-init
run-init:
	@ go mod tidy

.PHONY: run-dev
run-dev:
	@ go run main.go

.PHONY: run-test
run-test:
	@ go test ./... -count=1