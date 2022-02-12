FROM golang:1.13-stretch

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

SHELL ["/bin/bash", "-c"]
RUN apt update && apt-get install -y unzip

# install protc
WORKDIR /protoc
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
RUN unzip protoc-3.11.2-linux-x86_64.zip
RUN ln -s /protoc/bin/protoc /bin/protoc

# golang
WORKDIR /go/src/github.com/kikils/backend
COPY . .
RUN go get -u github.com/pilu/fresh
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
CMD ["fresh"]
