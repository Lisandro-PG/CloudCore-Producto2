FROM golang:1.23 AS build

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/app .

EXPOSE 8080

CMD ["./app"]
