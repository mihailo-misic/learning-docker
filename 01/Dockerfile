FROM golang:1.10

WORKDIR /go/src/app
COPY . .

RUN go get github.com/pilu/fresh
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["fresh"]
