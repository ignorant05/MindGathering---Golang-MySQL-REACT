FROM golang:alpine AS builder

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev

RUN CGO_ENABLED=0 GO111MODULE=on go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/db/migration ./db/migration

RUN apk add --no-cache ca-certificates mysql-client

EXPOSE 3000 

CMD ["./main"]
