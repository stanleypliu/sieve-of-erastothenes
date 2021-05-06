FROM golang:1.15

WORKDIR /go/src/sieve-of-eratosthenes

COPY . .

RUN go get 

CMD ["go", "run", "main.go"]