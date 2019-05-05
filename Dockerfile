FROM golang:1.12-alpine3.9

WORKDIR /go/src/github.com/src/taxio/gitcrow
ADD . /go/src/github.com/src/taxio/gitcrow

CMD ["go", "env"]
