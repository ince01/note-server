FROM golang:latest

WORKDIR /usr/note-server

RUN apt update && apt upgrade -y && \
  apt install -y git \
  make openssh-client

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
  && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air
