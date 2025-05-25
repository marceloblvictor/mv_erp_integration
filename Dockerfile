# build stage
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/web


# final stage
FROM alpine:latest

RUN adduser -D appuser
USER appuser

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

ENV COSMOS_ENDPOINT=https://mv-erp-integration-db.documents.azure.com:443

CMD ["./server"]
