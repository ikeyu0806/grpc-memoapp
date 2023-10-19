FROM golang:1.21-bullseye

WORKDIR /app

RUN apt-get update && apt-get install -y sqlite3

CMD ["go","run","/app/cmd/server/main.go"]
