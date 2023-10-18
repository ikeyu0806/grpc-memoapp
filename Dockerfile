FROM golang:1.21-bullseye

WORKDIR /app

CMD ["go","run","/app/cmd/server/main.go"]
