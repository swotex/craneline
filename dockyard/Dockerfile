# --- build stage ---
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

# --- runtime stage ---
FROM alpine:3.19
RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/api /bin/api

EXPOSE 8080
ENTRYPOINT ["/bin/api"]
