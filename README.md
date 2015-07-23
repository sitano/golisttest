# golisttest
Single list and naive Set implementation example

# How to run?

You have 3 ways of running the tests:

1) From linux shell expecting you have go installed and $GOPATH env set

```bash
go get github.com/sitano/golisttest
go test -v github.com/sitano/golisttest/... 
```

2) Using linux docker

```bash
git clone https://github.com/sitano/golisttest
docker build --rm -t golisttest ./golisttest/
docker run --rm -ti --name golisttest golisttest
```

3) Using vagrant from whatever OS you like

```
vagrant up
```
