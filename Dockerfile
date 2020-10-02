FROM golang:1.12.0-alpine3.9

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .
##RUN go build .

EXPOSE 3000

# Run the executable
CMD ["/app/main"]

