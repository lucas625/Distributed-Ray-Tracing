#!/bin/bash
source env_vars.sh

# Building images
docker build -t $DRT_TAG_PREFIX/drt-ray-tracing:$DRT_TAG_VERSION -f ray-tracing/Dockerfile ray-tracing
docker build -t $DRT_TAG_PREFIX/drt-ray-tracing-controller:$DRT_TAG_VERSION -f ray_tracing_controller/Dockerfile ray_tracing_controller

# Creating the network
docker network create drt-network

# Running the images
docker run --rm -d --name drt-ray-tracing-container --network=drt-network \
    -p 8081:8081 \
    $DRT_TAG_PREFIX/drt-ray-tracing:$DRT_TAG_VERSION
docker run --rm -d --name drt-ray-tracing-controller-container --network=drt-network \
    -p 8083:8083 \
    -e SECRET_KEY=$DRT_RAY_TRACING_CONTROLLER_SECRET_KEY \
    -e DEBUG=$DRT_RAY_TRACING_CONTROLLER_DEBUG \
    -e IMAGE_GENERATOR_ADDRESS=$DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS \
    -e RAY_TRACING_ADDRESS=$DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS \
    $DRT_TAG_PREFIX/drt-ray-tracing-controller:$DRT_TAG_VERSION
