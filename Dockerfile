ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

RUN apt update && \
    apt-get install -y cmake protobuf-compiler
WORKDIR /usr/src/app
COPY . .
RUN make deps_grpc
RUN make build


FROM debian:bookworm

COPY --from=builder /usr/src/app/target/mnstrapp-gossip /usr/local/bin/
CMD ["mnstrapp-gossip"]
