###########################
#1 Сборочный этап (builder)
###########################
FROM golang:1.24-alpine AS builder
LABEL stage=builder

#1.1 Рабочая директория внутри контейнера
WORKDIR /app

#1.2 Копируем файлы с зависимостями и скачиваем модули
COPY go.mod go.sum ./
RUN go mod download

#1.3 Копируем исходники приложения
COPY . .

#1.4 Собираем приложение
RUN go build -o calculator ./cmd/calculator

#########################
#2 Финальный легкий образ
#########################

#2.1
FROM alpine:latest

#2.2 Рабочая директория внутри контейнера
WORKDIR /app

#2.3 Копируем собранное приложение
COPY --from=builder /app/calculator ./
#2.3 Копируем файлы конфигурации
COPY --from=builder /app/configs ./configs

#2.4 Запуск
ENTRYPOINT [ "./calculator" ]