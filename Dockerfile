FROM golang:1.12.4

WORKDIR /go/src/github.com/nekonenene/quiz_app

RUN curl -fLo /usr/local/bin/air \
  https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air && \
  chmod +x /usr/local/bin/air

RUN curl -L https://github.com/k0kubun/sqldef/releases/download/v0.4.10/mysqldef_linux_arm64.tar.gz | \
  tar -xz -C /usr/local/bin && \
  chmod +x /usr/local/bin/mysqldef

COPY . .

CMD ["air"]
