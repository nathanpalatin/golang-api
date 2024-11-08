ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -v -o /run-app .

FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
