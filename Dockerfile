FROM golang:1.12.4

WORKDIR /go/src/github.com/nekonenene/quiz_app

RUN curl -fLo /usr/bin/air \
  https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air && \
  chmod +x /usr/bin/air

COPY . .

CMD ["/usr/bin/air"]
