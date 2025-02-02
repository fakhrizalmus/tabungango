FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]