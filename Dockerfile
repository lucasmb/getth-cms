# Dockerfile
FROM golang:1.22-alpine as build

WORKDIR /app
COPY . .

RUN apk update && apk add --no-cache make gcc libc-dev npm nodejs && apk -v cache clean &&\
    go install github.com/air-verse/air@latest &&\
    go install github.com/a-h/templ/cmd/templ@latest &&\
    go mod download && npm install

# FROM alpine:latest as run

# WORKDIR /app
# COPY --from=build /app/spendings ./run
# COPY --from=build /app/db.json ./db.json

# EXPOSE 8080

CMD ["air", "-c", ".air.toml"]