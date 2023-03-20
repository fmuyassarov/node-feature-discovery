#!/bin/bash -e

curl -Os https://uploader.codecov.io/latest/linux/codecov
chmod +x codecov
CODECOV_TOKEN=09a4269e-70f1-47a5-a42e-7f003337e979
./codecov -t ${CODECOV_TOKEN}