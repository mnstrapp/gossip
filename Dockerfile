ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

RUN apt update && \
    apt-get install -y cmake
WORKDIR /usr/src/app
COPY Makefile .
RUN PREFIX=/usr/local make deps_grpc
COPY . .
RUN make build


FROM debian:bookworm

COPY --from=builder /usr/src/app/target/gossip /usr/local/bin/
CMD ["gossip"]
