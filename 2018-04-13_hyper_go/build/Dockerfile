FROM alpine:3.6

ADD hello /go/bin/hello

EXPOSE 80

RUN apk add --no-cache ca-certificates

ENTRYPOINT /go/bin/hello
