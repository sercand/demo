FROM alpine:edge

RUN apk add --update ca-certificates nodejs libc6-compat && rm -rf /var/cache/apk/*

WORKDIR "/demo/feed"
ADD package.json /demo/feed/package.json
RUN npm install

ADD protos /protos
ADD server.js /demo/feed/server.js

EXPOSE 3020
CMD ["node","server.js"]
