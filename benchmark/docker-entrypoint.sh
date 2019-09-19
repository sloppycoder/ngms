#!/bin/sh

set -e
SLAVE_PROTO=${SLAVE_PROTO:-grpc}


./slave -proto $SLAVE_PROTO &
echo starting locust master
/usr/local/bin/locust --master
