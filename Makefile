
build:
	@go build -o bin/HttpHandler cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/HttpHandler

kill:
	@fuser -k 8080/tcp || true

