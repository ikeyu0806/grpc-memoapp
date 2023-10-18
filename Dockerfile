FROM golang:1.21-bullseye

RUN apt-get update && apt-get install -y \
    protobuf

WORKDIR /app

CMD ["go","run","/app/server/main.go"]
