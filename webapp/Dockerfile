FROM golang:1.13 AS builder
LABEL webapp=builder

ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app
COPY . .

RUN go get -d ./... && \
    go build \
    -a \
    -ldflags "-s -w -extldflags 'static'" \
    -installsuffix cgo \
    -tags netgo \
    -o /bin/web-app \
  .

FROM scratch
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /bin/web-app /
COPY --from=builder /go/src/app/templates /templates
USER 65534
CMD ["/web-app"]
