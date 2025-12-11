build:
	go build -o bin/cipolla ./cmd/cipolla

lint:
	go tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint run

test:
	go test -race -timeout 1h -coverprofile cp.out ./...

generate: build
	go generate ./...
