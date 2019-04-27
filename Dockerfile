FROM golang:1.12.4

WORKDIR /go/src/github.com/nekonenene/quiz_app

COPY . .

CMD ["go", "run", "api.go"]
