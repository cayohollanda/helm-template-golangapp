FROM golang:latest

WORKDIR /app

COPY . .

RUN ls
RUN go build -o main .

CMD ["./main"]
