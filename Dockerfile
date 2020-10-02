FROM golang:1.12.0-alpine3.9

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .
##RUN go build .

EXPOSE 8080

# Run the executable
CMD ["./main"]

