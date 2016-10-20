FROM alpine:3.4

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /demo
ADD index.html /demo/index.html
ADD public /demo/public
ADD frontend-linux-amd64 /demo/frontend
EXPOSE 3030

CMD ["/demo/frontend"]