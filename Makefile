build:
	@go build -o bin/jsonAPI.exe cmd/main.go
run: build
	@./bin/jsonAPI
test:
	@go test -v ./...