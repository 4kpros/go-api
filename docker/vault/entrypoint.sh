#!/bin/bash

# Copy config files
mkdir -p /etc/vault
mv /app/${APP_ENV}.vault.hcl /etc/vault/vault.hcl
find "/app/" -type f -name "*.vault.hcl" -delete

# Extract backup files
if [ -f "/app/vault-data.tar.gz" ]; then
    # Backup current data
    mkdir -p /vault/data/backups
    mv /vault/data/backups /app
    tar czvf /app/backups/$(date +%Y%m%d_%H%M%S).tar.gz /vault/data

    # Apply new backup
    rm -r /vault/data/*
    tar -zxvf /app/vault-data.tar.gz -C /vault/data
    rm /app/vault-data.tar.gz
    mv /app/backups /vault/data
fi

vault server -config=/etc/vault/vault.hcl
