FROM alpine:latest

ADD gin-framework /

EXPOSE 8080
CMD "./gin-framework"