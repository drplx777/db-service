FROM golang:1.24.4-alpine
WORKDIR /app

# Устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Собираем бинарь
RUN go build -o server ./cmd/server

# Запуск
EXPOSE 8000
ENTRYPOINT ["/app/server"]