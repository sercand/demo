FROM alpine:3.4

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

ADD feed-aggregator-linux-amd64 /demo/feed
ADD sample.json /demo/sample.json

EXPOSE 3010

CMD ["/demo/feed"]
