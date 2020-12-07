#!/bin/bash -e

source $(dirname "$0")/common.sh

pushd ${SERVER_DIR}
    gcloud builds submit --tag gcr.io/${PROJECT_ID}/server
popd