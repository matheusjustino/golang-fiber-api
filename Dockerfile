FROM golang:1.17.3-alpine3.15 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

FROM alpine:3.15

RUN apk add --no-cache bash

WORKDIR /app

COPY --from=build /app/server .

CMD ["./server"]