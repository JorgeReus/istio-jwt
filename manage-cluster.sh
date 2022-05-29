#!/bin/env bash

# Defaults
CLUSTER_NAME=istio-test
NUM_WORKER_NODES=3
REGISTRY_NAME=${CLUSTER_NAME}-registry
REGISTRY_PORT=5050

create_cluster() {
  k3d cluster create ${CLUSTER_NAME} -a ${NUM_WORKER_NODES} --registry-create ${REGISTRY_NAME}:0.0.0.0:${REGISTRY_PORT} --k3s-arg "--no-deploy=traefik@server:*"
}

destroy_cluster() {
  k3d cluster delete ${CLUSTER_NAME} 
}

push_image() {
  image=$1
  # Auth service
  docker tag ${image} localhost:${REGISTRY_PORT}/${image}
  docker push localhost:${REGISTRY_PORT}/${image}
  # Nginx
  # docker tag nginx:latest localhost:5050/nginx:v1.0
  # docker push localhost:5050/nginx:v1.0
}

ACTION=$1
shift
while getopts "n:a:" opt; do
  case "${opt}" in 
    n) CLUSTER_NAME=$OPTARG
      ;;
    a) NUM_WORKER_NODES=$OPTARG
      ;;
  esac
done     

case "$ACTION" in
 create)
  create_cluster
  ;;
 destroy)
  destroy_cluster
  ;;
 push)
  push_image $1
  ;;
 *)
  echo "Usage: (create|destroy) [-n <cluster-name>] [-a <worker-nodes-number>]"
  ;;
esac
