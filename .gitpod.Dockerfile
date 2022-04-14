FROM gitpod/workspace-go:latest

ENV GO_VERSION=1.18.1

RUN curl -fsSL https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz | tar xzs
