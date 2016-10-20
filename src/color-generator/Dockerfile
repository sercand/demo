FROM alpine:3.4

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

ADD color-generator-linux-amd64 /demo/color-generator

EXPOSE 4000

CMD ["/demo/color-generator"]