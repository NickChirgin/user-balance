FROM golang:1.20.0

WORKDIR /usr/src/avito

COPY . .
RUN go mod tidy