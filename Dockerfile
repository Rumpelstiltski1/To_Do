FROM golang:1.24-alpine AS builder

WORKDIR /app
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV ENV=production
RUN go build -o todo ./cmd/todo/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/todo .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
CMD migrate -path ./migrations -database "$DATABASE_URL" up && ./todo
