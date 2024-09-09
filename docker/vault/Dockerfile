FROM hashicorp/vault:latest AS vault

ENV LANG=C.UTF-8

WORKDIR /app

# Install packages
RUN apk update
RUN apk add --no-cache --upgrade bash

# Copy config files
COPY ./config/. ./
COPY ./tls ./

# Copy backup files
COPY ./backup/vault-data.tar.gz ./

# Copy entrypoint
COPY ./entrypoint.sh ./
RUN chmod +x ./entrypoint.sh

EXPOSE 8200

ENTRYPOINT ["/app/entrypoint.sh"]
