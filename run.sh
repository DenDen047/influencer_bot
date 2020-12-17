#!/bin/bash

IMG_NAME="denden047/influencer_bot"

docker build -t ${IMG_NAME} . && \
docker run -it --rm \
    -v "$PWD":/usr/src/myapp \
    -w /usr/src/myapp \
    ${IMG_NAME} \
    /bin/bash -c "go run main.go"
