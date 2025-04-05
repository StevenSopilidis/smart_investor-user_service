FROM golang:1.23-alpine AS Build

WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /user-service main.go

FROM alpine:latest

WORKDIR /
COPY --from=Build /user-service /user-service

ENTRYPOINT [ "/user-service" ]