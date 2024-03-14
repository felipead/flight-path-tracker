.PHONY: build
build:
	@go build -o bin/flight-path-tracker cmd/main.go

.PHONY: clean
clean:
	@rm -rf ./bin

.PHONY: ci
ci:
	@make lint
	@make test

.PHONY: lint
lint:
	@echo all good for now


.PHONY: test
test:
	go test -v ./...
