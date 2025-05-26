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
ENV AZURE_CLIENT_ID=383f12e1-9c1a-42be-bc97-432025916641
ENV AZURE_TENANT_ID=db9ea2b7-3e25-4cd2-ad77-be6d08ca0db2

CMD ["./server"]
