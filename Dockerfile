FROM golang:1.15.2-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 3000

# Run the executable
CMD ["/app/main"]

