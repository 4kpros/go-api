FROM golang:1.23.2 as builder

ENV LANG=C.UTF-8

WORKDIR /app

# Installing packages
RUN apt-get -y install bash make

COPY . .

# Build app
RUN make clean-cache
RUN make install
RUN make build

FROM alpine:latest as api

WORKDIR /app

# Copy app
COPY FROM builder ./docker/api/entrypoint.sh ./
RUN chmod +x ./app/.build/main

EXPOSE 9000

ENTRYPOINT ["make run"]
