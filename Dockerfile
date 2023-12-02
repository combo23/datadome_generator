FROM golang:1.20

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080 443

CMD ["./main"]
