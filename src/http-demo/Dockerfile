FROM alpine:3.4

RUN apk add --update ca-certificates nodejs && rm -rf /var/cache/apk/*

ADD index.js /demo/index.js

EXPOSE 3000

CMD ["node","/demo/index.js"]
