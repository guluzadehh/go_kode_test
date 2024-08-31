build:
	@go build -o bin/kode cmd/main.go

run: build
	@./bin/kode