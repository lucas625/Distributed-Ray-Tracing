FROM python:3.9.0-alpine3.12

# Install dependencies only
COPY requirements.txt /ray_tracing_controller/
RUN pip install -r /ray_tracing_controller/requirements.txt

# Setup the rest of the project to be executed
COPY . /ray_tracing_controller
WORKDIR /ray_tracing_controller

# Env vars
ENV SECRET_KEY ""
ENV DEBUG false
ENV IMAGE_GENERATOR_ADDRESS "http://127.0.0.1:8082/"
ENV RAY_TRACING_ADDRESS "http://127.0.0.1:8081/"

ENTRYPOINT ["./entrypoint.sh"]
