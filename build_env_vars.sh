#!/bin/bash

rm env_vars.sh

echo '#!/bin/bash' >> env_vars.sh

# General
echo 'export DRT_TAG_PREFIX="development"' >> env_vars.sh
echo 'export DRT_TAG_VERSION="1.0"' >> env_vars.sh
echo 'export DRT_IMAGE_PULL_POLICY="Never"' >> env_vars.sh

# DRT ray tracing controller
echo 'export DRT_RAY_TRACING_CONTROLLER_SECRET_KEY=' >> env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_DEBUG=false' >> env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_IMAGE_GENERATOR_ADDRESS=http://127.0.0.1:8082/' >> env_vars.sh
echo 'export DRT_RAY_TRACING_CONTROLLER_RAY_TRACING_ADDRESS=http://127.0.0.1:8081/' >> env_vars.sh

chmod +x env_vars.sh
