#!/bin/bash

if [ ! -f "config.tson" ]; then
    sed -i 's/MONGODB_HOST/'"$MONGODB_HOST"'/g' config.tson
    sed -i 's/MONGODB_PORT/'"$MONGODB_PORT"'/g' config.tson
    sed -i 's/MONGODB_NAME/'"$MONGODB_NAME"'/g' config.tson
    sed -i 's/MONGODB_USER/'"$MONGODB_USER"'/g' config.tson
    sed -i 's/MONGODB_PASS/'"$MONGODB_PASS"'/g' config.tson

    sed -i 's/JWT_SECRET/'"$JWT_SECRET"'/g' config.tson
    sed -i 's/MAIN_SECRET/'"$MAIN_SECRET"'/g' config.tson

    mv config.tson config.json
fi

exec "/app/lab-gate"