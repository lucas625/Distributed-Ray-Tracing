#!/bin/bash

rm docker_env_vars.sh

echo '#!/bin/bash' >> docker_env_vars.sh

# General
echo 'export DRT_TAG_PREFIX="gcr.io/distributed-ray-tracing"' >> docker_env_vars.sh
echo 'export DRT_TAG_VERSION="1.0"' >> docker_env_vars.sh
echo 'export DRT_IMAGE_PULL_POLICY="Always"' >> docker_env_vars.sh
echo 'export DRT_STATIC_IP="127.0.0.1:80"' >> docker_env_vars.sh

# DRT ray tracing controller
echo 'export DRT_RAY_TRACING_CONTROLLER_SECRET_KEY=test-secret-key-controller' >> docker_env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_DEBUG=false' >> docker_env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS=http://drt-image-generator:8082' >> docker_env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS=http://drt-ray-tracing:8081' >> docker_env_vars.sh

# DRT ray tracing
echo 'export DRT_RAY_TRACING_NUMBER_OF_THREADS=4' >> docker_env_vars.sh

# DRT image generator
echo 'export DRT_RAY_TRACING_IMAGE_GENERATOR_SECRET_KEY=test-secret-key-image-generator' >> docker_env_vars.sh
echo 'export DRT_RAY_TRACING_IMAGE_GENERATOR_DEBUG=false' >> docker_env_vars.sh

# DRT frontend
echo 'export DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL=$STATIC_IP' >> docker_env_vars.sh

chmod +x docker_env_vars.sh
