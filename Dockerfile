FROM golang:1.14.0-buster as builder

WORKDIR /workdir

COPY ./ /workdir/

RUN go build main.go

FROM debian:buster

COPY --from=builder /workdir/main /kubesu

ENTRYPOINT ["/kubesu"]
