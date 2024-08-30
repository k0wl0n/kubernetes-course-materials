FROM alpine:3.15.4

ADD docker-entrypoint.sh /

RUN chmod +x /docker-entrypoint.sh && apk update && apk add openssl

ENTRYPOINT ["/docker-entrypoint.sh"]
