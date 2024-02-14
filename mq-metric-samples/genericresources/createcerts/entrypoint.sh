#!/bin/bash

set -e

# Create the keystore
/opt/mqm/bin/runmqakm -keydb -create -db /opt/mqm/key.kdb -pw passw0rd -stash -type kdb

# Import certificate into the keystore
/opt/mqm/bin/runmqakm -cert  -import -target /opt/mqm/key.kdb -stashed -file /opt/mqm/server.pem

# Start the exporter
/opt/bin/mq_prometheus
