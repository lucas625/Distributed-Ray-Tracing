# Deploy on gcloud

- [Deploy on gcloud](#deploy-on-gcloud)
  - [Setup the environment variables](#setup-the-environment-variables)
  - [Point to correct project](#point-to-correct-project)
  - [Building the images](#building-the-images)
  - [Pushing the images to gcloud](#pushing-the-images-to-gcloud)
  - [Setup cluster](#setup-cluster)
  - [Create the static IP](#create-the-static-ip)
  - [Setup pods and services](#setup-pods-and-services)

## Setup the environment variables

```bash
# Build the environment variables.
./build_env_vars.sh

# Run the environment variables
# Remember to change the values as needed.
. ./env_vars.sh
```

## Point to correct project

```bash
# Performing login.
gcloud auth login
gcloud config set project distributed-ray-tracing
```

## Building the images

```bash
# Building the images.
docker build -t $DRT_TAG_PREFIX/drt-ray-tracing:$DRT_TAG_VERSION -f ray-tracing/Dockerfile ray-tracing
docker build -t $DRT_TAG_PREFIX/drt-ray-tracing-controller:$DRT_TAG_VERSION -f ray_tracing_controller/Dockerfile ray_tracing_controller
docker build -t $DRT_TAG_PREFIX/drt-image-generator:$DRT_TAG_VERSION -f image_generator/Dockerfile image_generator
docker build -t $DRT_TAG_PREFIX/drt-frontend:$DRT_TAG_VERSION --build-arg VUE_APP_RAY_TRACING_CONTROLLER_URL=$DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL -f frontend/Dockerfile frontend
docker build -t $DRT_TAG_PREFIX/drt-reverse-proxy:$DRT_TAG_VERSION -f reverse_proxy/Dockerfile reverse_proxy
```

## Pushing the images to gcloud

```bash
# Pushing images.
gcloud docker -- push $DRT_TAG_PREFIX/drt-ray-tracing:$DRT_TAG_VERSION
gcloud docker -- push $DRT_TAG_PREFIX/drt-ray-tracing-controller:$DRT_TAG_VERSION
gcloud docker -- push $DRT_TAG_PREFIX/drt-image-generator:$DRT_TAG_VERSION
gcloud docker -- push $DRT_TAG_PREFIX/drt-frontend:$DRT_TAG_VERSION
gcloud docker -- push $DRT_TAG_PREFIX/drt-reverse-proxy:$DRT_TAG_VERSION
```

## Setup cluster

```bash
gcloud login
# Creates the cluster.
# Remember to go to GKE and delete the cluster when it is no longer necessary.
gcloud container clusters create drt-e2-small --zone southamerica-east1-a --machine-type e2-small

# Update `kubectl` context in case the current context be not pointing to the cluster created.
kubectl config current-context
kubectl config get-contexts
kubectl config use-context [CONTEXT_NAME]

# Ensure connection to cluster.
gcloud container clusters get-credentials drt-e2-small --zone southamerica-east1-a
```

## Create the static IP

```bash
# Creates the static ip.
gcloud compute addresses create drt-static-ip --region southamerica-east1

# Add the ip listed by the command bellow to the environment variables.
gcloud compute addresses describe drt-static-ip --region southamerica-east1

# Set the variables with the static ip.
. ./env_vars.sh

# Delete the static IP when not using.
gcloud compute addresses delete drt-static-ip --region southamerica-east1
```

## Setup pods and services

```bash
# Run kubernetes.
cat kubernetes.yaml | sed \
    -e "s/\$\$DRT_TAG_PREFIX/$DRT_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$DRT_TAG_VERSION/$DRT_TAG_VERSION/" \
    -e "s/\$\$DRT_IMAGE_PULL_POLICY/$DRT_IMAGE_PULL_POLICY/" \
    -e "s/\$\$DRT_STATIC_IP/$DRT_STATIC_IP/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_SECRET_KEY/$DRT_RAY_TRACING_CONTROLLER_SECRET_KEY/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_DEBUG/$DRT_RAY_TRACING_CONTROLLER_DEBUG/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS/$DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS/$DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS/" \
    -e "s/\$\$DRT_RAY_TRACING_IMAGE_GENERATOR_SECRET_KEY/$DRT_RAY_TRACING_IMAGE_GENERATOR_SECRET_KEY/" \
    -e "s/\$\$DRT_RAY_TRACING_IMAGE_GENERATOR_DEBUG/$DRT_RAY_TRACING_IMAGE_GENERATOR_DEBUG/" \
    -e "s/\$\$DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL/$DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL/" | \
    kubectl apply -f -

# Deletes pods and services when no longer necessary.
cat kubernetes.yaml | sed \
    -e "s/\$\$DRT_TAG_PREFIX/$DRT_TAG_PREFIX_FOR_REPLACEMENT/" \
    -e "s/\$\$DRT_TAG_VERSION/$DRT_TAG_VERSION/" \
    -e "s/\$\$DRT_IMAGE_PULL_POLICY/$DRT_IMAGE_PULL_POLICY/" \
    -e "s/\$\$DRT_STATIC_IP/$DRT_STATIC_IP/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_SECRET_KEY/$DRT_RAY_TRACING_CONTROLLER_SECRET_KEY/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_DEBUG/$DRT_RAY_TRACING_CONTROLLER_DEBUG/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS/$DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS/" \
    -e "s/\$\$DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS/$DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS/" \
    -e "s/\$\$DRT_RAY_TRACING_IMAGE_GENERATOR_SECRET_KEY/$DRT_RAY_TRACING_IMAGE_GENERATOR_SECRET_KEY/" \
    -e "s/\$\$DRT_RAY_TRACING_IMAGE_GENERATOR_DEBUG/$DRT_RAY_TRACING_IMAGE_GENERATOR_DEBUG/" \
    -e "s/\$\$DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL/$DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL/" | \
    kubectl delete -f -
```
