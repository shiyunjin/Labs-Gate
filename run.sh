#!/bin/bash

if [ ! -f "config.json" ]; then
    sed -e "s/MONGODB_HOST/$MONGODB_HOST/" \
        -e "s/MONGODB_PORT/$MONGODB_PORT/" \
        -e "s/MONGODB_NAME/$MONGODB_NAME/" \
        -e "s/MONGODB_USER/$MONGODB_USER/" \
        -e "s/MONGODB_PASS/$MONGODB_PASS/" \
        -e "s/JWT_SECRET/$JWT_SECRET/" \
        -e "s/MAIN_SECRET/$MAIN_SECRET/" config.tson > config.json
fi

exec "/app/lab-gate"