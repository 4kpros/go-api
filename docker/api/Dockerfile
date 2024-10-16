# ----------------------- Stage building -----------------------

FROM golang:1.23.2 AS builder

ENV LANG=C.UTF-8

WORKDIR /app

# Install linux packages
RUN apt-get -y install bash make

COPY . .

# Install go packages
RUN make install
# Build app
RUN make build-container


# ----------------------- Stage serve -----------------------

FROM alpine:latest AS api

ENV LANG=C.UTF-8

WORKDIR /app

COPY --from=builder /app/.build/ ./
RUN chmod +x /app/main

EXPOSE 3100

ENTRYPOINT ["/app/main"]
