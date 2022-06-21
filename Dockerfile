FROM golang:1.18-alpine

RUN apk update

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go get -d ./...

RUN go install -v ./...

RUN go build -o binary

EXPOSE 2032

CMD ["/app/binary"]