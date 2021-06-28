#!/bin/sh -l

${{ inputs.list_of_image}} > test1.json

APP_NAME="${{ inputs.app_name }}"
JQ_ARGS=".[] | select(.spec.name == \"${APP_NAME}\") | .id"
APP_ID="$(doctl app list -ojson | jq -r "${JQ_ARGS}")"

doctl app get ${APP_ID} -ojson | yq eval -P - '[0].spec' > _temp.yaml
./main
doctl app update ${APP_ID} --spec _temp.yaml

echo "=> Deploying app ${APP_NAME} (${APP_ID})..."
doctl app create-deployment $APP_ID

# Wait for latest deployment to be active or failed.
while true; do
    PHASE="$(doctl app list-deployments ${APP_ID} -ojson | jq -r '.[0].phase')"
    echo "-- Phase=${PHASE}"
    if [ "${PHASE}" == "ACTIVE" ]; then
        echo "=> Success! 🎉"
        exit 0
    fi
    if [ "${PHASE}" == "FAILED" ]; then
        echo "=> Failure, check your app deploy logs. 🚨"
        exit 1
    fi
    sleep 3
done
