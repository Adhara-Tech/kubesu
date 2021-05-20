FROM golang:1.14.0-buster as builder

WORKDIR /workdir

COPY ./ /workdir/

RUN go build main.go

FROM debian:buster

COPY --from=builder /workdir/main /kubesu

RUN chmod +x /kubesu

RUN groupadd --gid 1000 adhara && useradd --uid 1000 --gid adhara --shell /bin/bash --create-home adhara

USER adhara

ENTRYPOINT ["/kubesu"]
