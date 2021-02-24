FROM golang:1.16.0-alpine3.13
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]