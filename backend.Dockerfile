FROM golang:1.17-alpine

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .
COPY ./.env .

RUN go mod download

COPY ./*.go ./

RUN go build -o backend

ENTRYPOINT [ "./backend" ]
