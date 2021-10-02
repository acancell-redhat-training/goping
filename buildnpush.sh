#!/usr/bin/env bash

if [ "$DEBUG" = 1 ]; then
  set -x
fi

now="$(date -u +%Y%m%dT%H%M%SZ)"

echo "Build image"
podman build \
  -t goping:$now -t goping:latest \
  -t quay.io/acancell-redhat-training/goping:$now -t quay.io/acancell-redhat-training/goping:latest \
  .

echo "Push image"
podman push quay.io/acancell-redhat-training/goping:$now
podman push quay.io/acancell-redhat-training/goping:latest
