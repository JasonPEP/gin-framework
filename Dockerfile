FROM alpine:latest

ADD main /

EXPOSE 8080
CMD "./main"