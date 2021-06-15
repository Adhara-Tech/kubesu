FROM golang:1.14.0-buster as builder

WORKDIR /workdir

COPY ./go.mod /workdir/go.mod
COPY ./go.sum /workdir/go.sum

RUN go mod download

COPY ./ /workdir/

RUN go build main.go

FROM debian:buster

COPY --from=builder /workdir/main /kubesu

RUN chmod +x /kubesu

RUN groupadd --gid 1000 kubesu && useradd --uid 1000 --gid kubesu --shell /bin/bash --create-home kubesu

USER kubesu

ENTRYPOINT ["/kubesu"]
