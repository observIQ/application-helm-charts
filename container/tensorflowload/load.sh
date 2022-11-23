#!/bin/sh

while :
do
    curl -v \
        -d '{"instances": [13.0, 23.0, 53.0, 103.0, 503.0, 15, 17]}' \
        -X POST \
        http://localhost:8501/v1/models/half_plus_two:predict
    sleep 10
done
