#!/bin/bash -e

readonly CURDIR=$(cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd)
readonly SERVER_DIR=$(cd ${CURDIR}/..; pwd)
readonly PROJECT_ID="go-gql-pg-poc"
readonly PROJECT_GCP_REGION="asia-south1"