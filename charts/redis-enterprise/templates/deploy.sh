#!/usr/bin/env bash

set -e

kubectl apply -f namespace.yaml

kubectl apply -f operator-latest.yaml -n redis-0
kubectl apply -f operator-latest.yaml -n redis-1

kubectl apply -f cluster.yaml