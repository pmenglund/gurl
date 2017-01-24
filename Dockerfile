FROM alpine
MAINTAINER martin@englund.nu

RUN apk add --update openssl
RUN wget -q -O /gurl https://github.com/pmenglund/gurl/releases/download/pre-release/linux_amd64_gurl
RUN chmod 755 /gurl

HEALTHCHECK --interval=15s --timeout=5s CMD /gurl http://localhost:8080/

COPY server /server

CMD ["/server"]
