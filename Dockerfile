FROM golang:1.12

ENV GIN_MODE release

EXPOSE 443

EXPOSE 80

# install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN go get -u github.com/kalderasoft/go-auth

WORKDIR src/github.com/kalderasoft/go-auth

RUN dep ensure

CMD ["go", "run", "main/main.go"]