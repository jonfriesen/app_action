#!/bin/sh -l

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
