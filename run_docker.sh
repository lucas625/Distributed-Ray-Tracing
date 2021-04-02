#!/bin/bash
source docker_env_vars.sh

# Building images
docker build -t $DRT_TAG_PREFIX/drt-ray-tracing:$DRT_TAG_VERSION -f ray-tracing/Dockerfile ray-tracing
docker build -t $DRT_TAG_PREFIX/drt-ray-tracing-controller:$DRT_TAG_VERSION -f ray_tracing_controller/Dockerfile ray_tracing_controller
docker build -t $DRT_TAG_PREFIX/drt-image-generator:$DRT_TAG_VERSION -f image_generator/Dockerfile image_generator
docker build -t $DRT_TAG_PREFIX/drt-frontend:$DRT_TAG_VERSION --build-arg VUE_APP_RAY_TRACING_CONTROLLER_URL=$DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL -f frontend/Dockerfile frontend
docker build -t $DRT_TAG_PREFIX/drt-reverse-proxy:$DRT_TAG_VERSION -f reverse_proxy/Dockerfile reverse_proxy


# Creating the network
docker network create drt-network

# Removing old containers
docker rm -f drt-ray-tracing
docker rm -f drt-image-generator
docker rm -f drt-ray-tracing-controller
docker rm -f drt-frontend
docker rm -f drt-reverse-proxy

# Running the images
docker run --rm -d --name drt-ray-tracing --network=drt-network \
    -e NUMBER_OF_THREADS=$DRT_RAY_TRACING_NUMBER_OF_THREADS \
    $DRT_TAG_PREFIX/drt-ray-tracing:$DRT_TAG_VERSION
docker run --rm -d --name drt-image-generator --network=drt-network \
    -e SECRET_KEY=$DRT_RAY_TRACING_IMAGE_GENERATOR_SECRET_KEY \
    -e DEBUG=$DRT_RAY_TRACING_IMAGE_GENERATOR_DEBUG \
    $DRT_TAG_PREFIX/drt-image-generator:$DRT_TAG_VERSION
docker run --rm -d --name drt-ray-tracing-controller --network=drt-network \
    -e SECRET_KEY=$DRT_RAY_TRACING_CONTROLLER_SECRET_KEY \
    -e DEBUG=$DRT_RAY_TRACING_CONTROLLER_DEBUG \
    -e IMAGE_GENERATOR_ADDRESS=$DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS \
    -e RAY_TRACING_ADDRESS=$DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS \
    $DRT_TAG_PREFIX/drt-ray-tracing-controller:$DRT_TAG_VERSION
docker run --rm -d --name drt-frontend --network=drt-network \
    $DRT_TAG_PREFIX/drt-frontend:$DRT_TAG_VERSION
docker run --rm -d --name drt-reverse-proxy --network=drt-network \
    -p 80:80 \
    $DRT_TAG_PREFIX/drt-reverse-proxy:$DRT_TAG_VERSION
