#!/bin/sh

set -e
SLAVE_PROTO=${SLAVE_PROTO:-grpc}

echo starting slave with protocol $SLAVE_PROTO
./slave -proto $SLAVE_PROTO &

echo starting locust master
/usr/local/bin/locust --master

