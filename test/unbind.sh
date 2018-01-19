#!/bin/bash -e

INSTANCE_ID=$1
BINDING_ID=$2
PLAN_UUID="7f4a5e35e4af2beb70076e72fab0b7ff"

if [ "$INSTANCE_ID" = "" ]
then
  echo "Usage: $0 <instance uuid> <binding uuid>"
  exit
fi

if [ "$BINDING_ID" = "" ]
then
  echo "Usage: $0 <instance uuid> <binding uuid>"
  exit
fi

curl \
    -k \
    -X DELETE \
    -H "Authorization: bearer $(oc whoami -t)" \
    -H "Content-type: application/json" \
    -H "Accept: application/json" \
    -H "X-Broker-API-Originating-Identity: " \
    "https://asb-1338-ansible-service-broker.172.17.0.1.nip.io/ansible-service-broker/v2/service_instances/$INSTANCE_ID/service_bindings/$BINDING_ID?accepts_incomplete=true&plan_id=$PLAN_UUID"


