build:
	@go build -o bin/weather-api cmd/main.go

run: build
	@./bin/weather-api