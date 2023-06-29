FROM golang:1.19.0-alpine as builder

RUN apk add --update build-base

WORKDIR /workdir

COPY ./go.mod /workdir/go.mod
COPY ./go.sum /workdir/go.sum

RUN go mod download

COPY ./ /workdir/

RUN go build main.go

FROM alpine:3.18.2

COPY --from=builder /workdir/main /kubesu

RUN chmod +x /kubesu

RUN addgroup -g 1000 -S kubesu && adduser -u 1000 -G kubesu -S kubesu

USER kubesu

ENTRYPOINT ["/kubesu"]
