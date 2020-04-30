FROM debian:buster-slim

EXPOSE 9090

WORKDIR /

ADD bin/headers /headers

CMD ["/headers"]
