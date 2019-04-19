FROM erkanzileli/golang1.12:dep0.5.1

ENV GIN_MODE release

EXPOSE 443

EXPOSE 8000

RUN go get -u github.com/kalderasoft/go-auth

WORKDIR src/github.com/kalderasoft/go-auth

RUN dep ensure

CMD ["go", "run", "main/main.go"]