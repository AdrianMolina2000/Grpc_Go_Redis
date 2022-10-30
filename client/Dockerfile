FROM golang:1.16.5-alpine

WORKDIR /
COPY . .

RUN go mod download

EXPOSE 50051
EXPOSE 2000

CMD ["go", "run", "client_api.go"]