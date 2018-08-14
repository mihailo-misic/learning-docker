FROM golang:1.10

WORKDIR /go/src/github.com/mihailo-misic/learning-docker
COPY . .

RUN go get github.com/pilu/fresh

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT fresh
