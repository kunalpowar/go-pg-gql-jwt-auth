#!/bin/bash -e

source $(dirname "$0")/common.sh

gcloud run deploy server --image gcr.io/go-gql-pg-poc/server --platform managed
  --add-cloudsql-instances "${PROJECT_ID}:${PROJECT_GCP_REGION}:server" \
  --set-env-vars INSTANCE_CONNECTION_NAME="${PROJECT_ID}:${PROJECT_GCP_REGION}:server" \
  --set-env-vars DB_USER="postgres" \
  --set-env-vars DB_PASS='yfwu~sVA"V9_>~(U' \
  --set-env-vars DB_NAME="gopggqlauth"