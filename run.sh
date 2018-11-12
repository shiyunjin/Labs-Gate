#!/usr/bin/env bash

if [ ! -f "config.json" ]; then
    if [ -z "$MONGODB_HOST" ]; then
        echo "MONGODB_HOST is null"
        exit 1
    fi
    if [ -z "$MONGODB_PORT" ]; then
        echo "MONGODB_PORT is null"
        exit 1
    fi
    if [ -z "$MONGODB_NAME" ]; then
        echo "MONGODB_NAME is null"
        exit 1
    fi
    if [ -z "$JWT_SECRET" ]; then
        echo "JWT_SECRET is null"
        exit 1
    fi
    if [ -z "$MAIN_SECRET" ]; then
        echo "MAIN_SECRET is null"
        exit 1
    fi
    sed -e "s/MONGODB_HOST/$MONGODB_HOST/g" \
        -e "s/MONGODB_PORT/$MONGODB_PORT/g" \
        -e "s/MONGODB_NAME/$MONGODB_NAME/g" \
        -e "s/MONGODB_USER/$MONGODB_USER/g" \
        -e "s/MONGODB_PASS/$MONGODB_PASS/g" \
        -e "s/JWT_SECRET/$JWT_SECRET/g" \
        -e "s/MAIN_SECRET/$MAIN_SECRET/g" config.tson > config.json

    chmod 777 config.json
fi

exec "/app/lab-gate"