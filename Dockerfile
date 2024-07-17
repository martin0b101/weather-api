# Use the official Golang image to create a build artifact
FROM golang:1.22.5 

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./weather-api ./cmd/main.go

CMD ["./weather-api"]
