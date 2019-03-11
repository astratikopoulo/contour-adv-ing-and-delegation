FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM golang:1.12.1 AS build
WORKDIR /serve

COPY mountains.go mountains.go
RUN CGO_ENABLED=0 GOOS=linux GOFLAGS=-ldflags=-w go build -o /go/bin/serve -ldflags=-s -v mountains.go

FROM scratch AS final
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/bin/serve /bin/serve