This directory contains scripts as examples of building and running
the various collector programs. The files are:

* `buildMonitors.sh`: Creates a container and uses it as the build environment
to build all the collectors. Output goes to a real local directory, by default `$HOME/tmp/mq-metric-samples`,
which holds the binaries, the default YAML configuration files and MQSC scripts that can be
used to run the collectors as MQ Services. Depending on your system configuration, you might need
to add `--privileged` to the `docker run` command, or set the OUTDIR environment variable to point at a 
different directory with suitable permissions.  

* `buildInDocker.sh`: Used inside the container created by the previous script. It downloads the
MQ Redistributable Client needed for compiling the packages, and then does the `go build`
operations.

* `buildRuntime.sh`: Create a container that holds the runtime for a particular collector. By
default, build containers for all collectors, but can give a command line list to explicitly ask
for a subset. For example, `buildRuntime.sh mq_prometheus mq_json`. The container is tagged, ready to run.

* `runMonitor.sh`: Execute one of the containers build in the previous step. The configuration file
is mounted from the host image. This is the file you'll most likely want to change, to match your
own setup. But it's a useful model to show how parameters can be passed.

* `runMonitorTLS.sh`: Execute the Prometheus monitor, with a CCDT and keystore configuration for TLS 
connections to the queue manager. 

* `buildMonitors.bat`: Example of building on Windows. It assumes you
already have the MQ client package installed, and have a suitable `gcc`
compiler available under c:\TDM-GCC-64. Output of the binaries and
YAML files go into `%GOPATH%\bin`.

* `docker-compose.yaml`: A fragment of what might go into a larger composition file. In its current
state it tries to pull an image from a repository even though it has been built and exists locally.

* `buildBuildah.sh`: An alternative approach to building a runnable container image. It uses `buildah`, `podman`
and the Red Hat UBI containers. It uses a multi-stage model, one to compile the code and a second phase that
uses a slimmer runtime image.
