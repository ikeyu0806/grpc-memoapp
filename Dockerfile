FROM golang:1.21-bullseye

WORKDIR /app

RUN apt-get update && apt-get install -y sqlite3
RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD ["go","run","/app/cmd/server/main.go"]
