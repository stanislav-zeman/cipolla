.PHONY: build
build:
	go build -o bin/cipolla ./cmd/cipolla

.PHONY: lint
lint:
	go tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint run

.PHONY: test
test:
	go test -race -timeout 1h -coverprofile cp.out ./...

.PHONY: generate
generate: clean build
	go generate ./...

.PHONY: clean
clean:
	rm -rf bin example
