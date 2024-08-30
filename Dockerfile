FROM golang:latest

WORKDIR /app

COPY ./app .
RUN go mod tidy && go build -o /app/main

CMD ["/app/main"]
