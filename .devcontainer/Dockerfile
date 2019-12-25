FROM ubuntu:latest

# Minimum required packages
RUN apt-get update && apt-get install -y \
    wget \
    git \
    vim

# golang
ARG GO_VERSION=1.13.5
ARG OS=linux
ARG ARCH=amd64
RUN wget -O ~/go${GO_VERSION}.${OS}-${ARCH}.tar.gz https://storage.googleapis.com/golang/go${GO_VERSION}.${OS}-${ARCH}.tar.gz \
 && tar -C /usr/local -xzf ~/go${GO_VERSION}.${OS}-${ARCH}.tar.gz \
 && rm -rf ~/go${GO_VERSION}.${OS}-${ARCH}.tar.gz \
 && echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc \
 && echo "export GOPATH=~/go" >> ~/.bashrc
