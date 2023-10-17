FROM golang:1.21-bullseye

WORKDIR /app

CMD ["go","run","/app/server/main.go"]
