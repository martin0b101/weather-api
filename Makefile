build:
	@go build -o bin/weather-api cmd/main.go

test:
	@go test ./...

run: build
	@./bin/weather-api