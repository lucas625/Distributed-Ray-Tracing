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

ENTRYPOINT ["./entrypoint.sh"]
