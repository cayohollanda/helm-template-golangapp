FROM golang:latest

WORKDIR /app

COPY ./src/ /app/src/

RUN go build -o main /app/src/

CMD ["./main"]
