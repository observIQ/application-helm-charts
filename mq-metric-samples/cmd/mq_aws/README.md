# MQ Exporter for Amazon CloudWatch monitoring

This README should be read in conjunction with the repository-wide
[README](https://github.com/ibm-messaging/mq-metric-samples/blob/master/README.md) 
that covers features common to all of the collectors in this repository.

This directory contains the code for a monitoring solution
that exports queue manager data to a CloudWatch data collection
system. It also contains configuration files to run the monitor program

An example Grafana dashboard is included, to show how queries might
be constructed. The data shown is the same as in the corresponding
Prometheus-based dashboard, also in this repository.
To use the dashboard,
create a data source in Grafana called "MQ CloudWatch" that points at your
AWS server, and then import the JSON file. Grafana does not make it as easy
to handle wildcard queries to CloudWatch, so this dashboard explicitly names
the queues to monitor. There may be better solutions using templates, but
that starts to get more complex than I want to show in this example.

## Configuring MQ
It is convenient to run the monitor program as a queue manager service.
This directory contains an MQSC script to define the service. In fact, the
service definition points at a simple script which sets up any
necessary environment and builds the command line parameters for the
real monitor program. As the last line of the script is "exec", the
process id of the script is inherited by the monitor program, and the
queue manager can then check on the status, and can drive a suitable
`STOP SERVICE` operation during queue manager shutdown.

Edit the MQSC script and the shell script to point at appropriate directories
where the program exists, and where you want to put stdout/stderr.
Ensure that the ID running the queue manager has permission to access
the programs and output files.

There are a number of required parameters to configure the service, including
the queue manager name, how to reach a database, and the frequency of reading
the queue manager publications. Look at the mq_aws.sh script or the sample yaml
file to see how to provide these parameters.

For authentication, the
program is expecting to pick up the keys from $HOME/.aws/credentials. That
file is not explicitly referenced in this code; it's used automatically by
the AWS toolkit. You may need to provide a region
name as a command-line option if it is not mentioned in the credentials
file.

The queue manager will usually generate its publications every 10 seconds. However, the
default interval being used by this monitor program is set to 60 seconds for reading those
publications because of the slower rate that CloudWatch monitoring usually works at, and
to reduce the number of calls to CloudWatch which do contribute to AWS charges.

## Configuring CloudWatch
No special configuration is required for CloudWatch. You will see the MQ
data appear in the Custom Metrics options, with a default namespace of "IBM/MQ".
There are two sets of metrics in that namespace, one for the queue manager
(with the qmgr filter) and one for the queues (with the object,qmgr filter).

## Metrics
Once the monitor program has been started,
you will see metrics being available.
console. All of the queue
manager values are given a tag of the queue manager name; all of the other object-based values
are tagged with both the object and queue manager names.

The example Grafana dashboard shows how queries can be constructed to extract data
about specific queues or the queue manager.

More information on the metrics collected through the publish/subscribe
interface can be found in the [MQ KnowledgeCenter](https://www.ibm.com/docs/en/ibm-mq/latest?topic=trace-metrics-published-system-topics)
with further description in [an MQDev blog entry](https://community.ibm.com/community/user/integration/viewdocument/statistics-published-to-the-system?CommunityKey=183ec850-4947-49c8-9a2e-8e7c7fc46c64&tab=librarydocuments)

The metrics stored in the database are named after the
descriptions that you can see when running the amqsrua sample program, but with some
minor modifications to match a more useful style.
