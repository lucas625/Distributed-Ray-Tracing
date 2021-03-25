#!/bin/bash

rm env_vars.sh

echo '#!/bin/bash' >> env_vars.sh

# General
echo 'export DRT_TAG_PREFIX="distributed-ray-tracing"' >> env_vars.sh
echo 'export DRT_TAG_VERSION="1.0"' >> env_vars.sh
echo 'export DRT_IMAGE_PULL_POLICY="Always"' >> env_vars.sh
echo 'export DRT_STATIC_IP="http://127.0.0.1:80"' >> env_vars.sh

# DRT ray tracing controller
echo 'export DRT_RAY_TRACING_CONTROLLER_SECRET_KEY=test-secret-key-controller' >> env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_DEBUG=false' >> env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS=http://drt-image-generator:8082' >> env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS=http://drt-ray-tracing:8081' >> env_vars.sh

# DRT image generator
echo 'export DRT_RAY_TRACING_IMAGE_GENERATOR_SECRET_KEY=test-secret-key-image-generator' >> env_vars.sh
echo 'export DRT_RAY_TRACING_IMAGE_GENERATOR_DEBUG=false' >> env_vars.sh

# DRT frontend
echo 'export DRT_FRONTEND_VUE_APP_RAY_TRACING_CONTROLLER_URL=$STATIC_IP' >> env_vars.sh

chmod +x env_vars.sh
