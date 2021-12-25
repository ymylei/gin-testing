FROM ubuntu:latest

RUN ["apt-get", "update"] 
RUN ["apt-get", "upgrade", "-y"] 
RUN ["apt-get", "install", "ca-certificates", "-y"]

RUN ["mkdir", "-p", "/webserver"]
COPY ["./out/gin-testing-linux-arm64", "/webserver"]

ENV GIN_MODE=release
EXPOSE 8080
CMD ["/bin/bash", "/webserver/gin-testing-linux-arm64"]
