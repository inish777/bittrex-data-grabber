FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go-wrapper download

CMD ["go-wrapper", "run"]