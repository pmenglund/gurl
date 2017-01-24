FROM alpine
MAINTAINER martin@englund.nu

RUN apk add --update openssl
RUN wget -q -O /gurl https://github.com/pmenglund/gurl/releases/download/pre-release/linux_amd64_gurl
# TODO: for some reason gox builds a dynamically linked binary, so this hackery is needed in alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN chmod 755 /gurl

HEALTHCHECK --interval=15s --timeout=5s CMD gurl http://localhost:8080/

COPY server /server

CMD ["/server"]
