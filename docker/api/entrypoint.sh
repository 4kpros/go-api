#!/bin/bash

# Get .env file
curl ${VAULT_ADDR}/v1/snip/data/api/env/${APP_ENV} \
  --request GET \
  --header "X-Vault-Token: ${VAULT_TOKEN}" \
  | jq -r '.data.data.env' > ./.env
envFilePath="./.env"
if [ ! -f "$envFilePath" ]; then
  echo "Missing ${envFilePath} !"
  exit 1
fi

# Check if the .env file is not empty
read -r LINE < "$envFilePath"
# Remove leading and trailing whitespaces, and carriage return
CLEANED_LINE=$(echo "$LINE" | awk '{$1=$1};1' | tr -d '\r')
if [ -z "$CLEANED_LINE" ]; then
  echo "Missing ${envFilePath} !"
  exit 1
fi

# Get jwt private.pem and public.pem files
mkdir -p ./config/jwt
curl ${VAULT_ADDR}/v1/snip/data/api/jwt/${APP_ENV} \
  --request GET \
  --header "X-Vault-Token: ${VAULT_TOKEN}" \
  | jq -r '.data.data.private' > ./config/jwt/private.pem
curl ${VAULT_ADDR}/v1/snip/data/api/jwt/${APP_ENV} \
  --request GET \
  --header "X-Vault-Token: ${VAULT_TOKEN}" \
  | jq -r '.data.data.public' > ./config/jwt/public.pem
jwtFilePath1="./config/jwt/private.pem"
jwtFilePath2="./config/jwt/private.pem"
if [ ! -f "$jwtFilePath1" ]; then
  echo "Missing ${jwtFilePath1} !"
  exit 1
fi
if [ ! -f "$jwtFilePath2" ]; then
  echo "Missing ${jwtFilePath2} !"
  exit 1
fi

# Create database
symfony console do:da:cr

# Create migrations from entities
symfony console ma:mi

# Load migrations
symfony console do:mi:mi

# Load fixures
symfony console doctrine:fixtures:load

# Enable tls support (only on dev, staging and prod)
if [[ "$APP_ENV" == "prod" || "$APP_ENV" == "staging" ]]; then
  symfony server:ca:install
fi

# Start the server
symfony server:start --port=9000
