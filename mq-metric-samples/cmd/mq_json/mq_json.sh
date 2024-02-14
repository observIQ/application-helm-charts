#!/bin/bash

# This is used to start an IBM MQ monitoring service for JSON-formatted output

# The queue manager name comes in from the MQ service configuration as
# the only command line parameter.
qMgr=$1

# Set the environment to ensure we pick up libmqm.so etc
# This assumes there is an MQ installation in the default location, even
# if it is not the one associated with the queue manager
# If this is a client connection, then deal with no known qmgr of the given name.
. /opt/mqm/bin/setmqenv -m $qMgr -k >/dev/null 2>&1
if [ $? -ne 0 ]
then
  . /opt/mqm/bin/setmqenv -s -k
fi
# A list of queues to be monitored is given here.
# It is a set of names or patterns ('*' only at the end, to match how MQ works),
# separated by commas. When no queues match a pattern, it is reported but
# is not fatal.
queues="APP.*,MYQ.*"

# An alternative is to have a file containing the patterns, and named
# via the ibmmq.monitoredQueuesFile option.

# Do similar for channels
channels="TO.*,SYSTEM.DEF.SVRCONN"

# And other parameters that may be needed
# See config.go for all recognised flags
interval="5s"

ARGS="-ibmmq.queueManager=$qMgr"
ARGS="$ARGS -ibmmq.monitoredQueues=$queues"
ARGS="$ARGS -ibmmq.monitoredChannels=$channels"
ARGS="$ARGS -ibmmq.monitoredTopics=#"
ARGS="$ARGS -ibmmq.monitoredSubscriptions=*"
ARGS="$ARGS -ibmmq.interval=$interval"
ARGS="$ARGS -ibmmq.useStatus=true"
ARGS="$ARGS -log.level=error"

# Start via "exec" so the pid remains the same.
#
# This program will send all log info to stderr
#
# Change this line to match wherever you have installed the MQ monitor program
# You probably also want to do something with the stdout from the program,
# such as sending it to a monitoring solution that understands the format.

exec /usr/local/bin/mqgo/mq_json $ARGS
