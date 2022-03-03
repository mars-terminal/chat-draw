FROM golang:1.17.7-alpine3.15

ENV GOPATH=/

COPY ./ ./

RUN go build -o main .

CMD ["./main"]
