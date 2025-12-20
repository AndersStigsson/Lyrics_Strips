FROM golang:1.25-alpine

WORKDIR /app

COPY ./backend/ .
COPY ./backend/.env .

RUN go mod download

RUN go build -o backend ./cmd/server

ENTRYPOINT [ "./backend" ]
