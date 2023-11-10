#!/bin/bash

# Create a config
influx config create --config-name sample-config --host-url http://influxdb-sample-app:80 --org influxdata --token nrlukwAn13XLMIlYeDTg7g47T28cEG4P --active

# Create buckets
influx bucket create --name loadgen-0
influx bucket create --name loadgen-1

# Write data to buckets
influx write --bucket loadgen-0 --file /influxdb2-sample-data-master/air-sensor-data/air-sensor-data.lp
influx write --bucket loadgen-1 --file /influxdb2-sample-data-master/bitcoin-price-data/currentprice.lp
influx write --bucket loadgen-0 --file /influxdb2-sample-data-master/noaa-ndbc-data/latest-observations.lp
influx write --bucket loadgen-1 --file /influxdb2-sample-data-master/bird-migration-data/bird-migration.line
influx write --bucket loadgen-0 --file /influxdb2-sample-data-master/usgs-earthquake-data/all_week.lp

# Query data from buckets for 5m
end=$((SECONDS+300))
while [ $SECONDS -lt $end ]
do
    influx query 'from(bucket:"loadgen-0") |> range(start:-1h)'
    influx query 'from(bucket:"loadgen-1") |> range(start:-1h)'
done

# Delete buckets
influx bucket delete --name loadgen-0
influx bucket delete --name loadgen-1
