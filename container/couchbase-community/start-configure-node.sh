#!/bin/bash

set -e

# Start couchbase server, send process to background
/entrypoint.sh couchbase-server &

# Configure couchbase server
/base-config-script.sh
