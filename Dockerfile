FROM golang:1.11

WORKDIR /go/src/github.com/metapods/metapods
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT /go/bin/metapods

EXPOSE 8080