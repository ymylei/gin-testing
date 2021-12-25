FROM ubuntu:latest

RUN ["apt-get", "update"] 
RUN ["apt-get", "upgrade", "-y"] 
RUN ["apt-get", "install", "ca-certificates", "-y"]

RUN ["mkdir", "-p", "/webserver"]
COPY ["./out/gin-testing-linux-*", "/webserver"]
COPY ["./scripts/server-start.sh", "/webserver"]

ENV GIN_MODE=release
EXPOSE 8080
