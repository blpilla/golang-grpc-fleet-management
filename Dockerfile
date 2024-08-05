# Etapa de construção
FROM golang:1.22.5-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download

COPY . ./

RUN go build -o main ./cmd/server

# Instalar o migrate no estágio de construção
RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    chmod +x /usr/local/bin/migrate

# Etapa de execução
FROM alpine:latest

# Instalar Go e outras dependências no estágio final
RUN apk add --no-cache go bash

WORKDIR /root/

COPY --from=build /app/main .
COPY --from=build /usr/local/bin/migrate /usr/local/bin/migrate
COPY --from=build /app .
COPY wait-for-it.sh /root/wait-for-it.sh

RUN chmod +x /root/wait-for-it.sh

ENTRYPOINT ["./wait-for-it.sh", "db", "5432", "--", "./main"]
