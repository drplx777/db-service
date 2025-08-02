FROM golang:1.24.4-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -o /server ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /server /app/server
EXPOSE 8000
ENTRYPOINT ["/app/server"]
