ARG GO_VERSION
FROM golang:${GO_VERSION}
ARG PROTOBUF_VERSION=3.19.4

RUN apt-get update && apt-get install unzip

WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip
RUN unzip protoc.zip
RUN mv bin/* /usr/local/bin/
RUN mv include/* /usr/local/include/

WORKDIR /usr/src/go-con-2022-spring-sample
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY pkg/domain/proto/server/options pkg/domain/proto/server/options
COPY cmd/protoc-gen-sample cmd/protoc-gen-sample
RUN go install ./cmd/protoc-gen-sample
COPY cmd/gen-file-remover cmd/gen-file-remover
RUN go install ./cmd/gen-file-remover

WORKDIR ${GOPATH}

ENTRYPOINT ["/usr/local/bin/protoc"]
