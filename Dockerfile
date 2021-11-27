FROM golang:1.17

WORKDIR /go/src/github.com/httpcutapi

COPY . .
 
RUN go get -u github.com/lib/pq

RUN go build -o main .

CMD ["./main"]
