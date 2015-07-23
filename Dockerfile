FROM golang

ADD . /go/src/github.com/sitano/golisttest

CMD ["go", "test", "-v", "github.com/sitano/golisttest/..."]
