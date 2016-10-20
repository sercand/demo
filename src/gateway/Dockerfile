FROM alpine:3.4

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

ADD gateway-linux-amd64 /demo/gateway

EXPOSE 8080

CMD ["/demo/gateway"]
