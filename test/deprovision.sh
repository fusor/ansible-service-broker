#!/bin/bash -e

instanceUUID="688eea24-9cf9-43e3-9942-d1863b2a16af"

curl \
  -X DELETE \
  -H 'X-Broker-API-Version: 2.9' \
  -v \
  http://localhost:1338/v2/service_instances/$instanceUUID
