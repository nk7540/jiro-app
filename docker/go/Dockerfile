FROM golang:1.16

RUN apt update
RUN apt install -y protobuf-compiler curl vim
RUN curl -ysL install-node.now.sh/lts | bash

WORKDIR /go/artics-api

ADD . /go/artics-api

ENV MYSQL_PASSWORD password
RUN export GCP_SERVICE_KEY_JSON="$(< ../../credentials/firebase/settings.json)" 

RUN go mod download
RUN go install golang.org/x/tools/gopls@latest
RUN go install golang.org/x/tools/cmd/goimports
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
RUN go install github.com/rubenv/sql-migrate/sql-migrate@latest
RUN go install github.com/pilu/fresh@latest

RUN sql-migrate up
