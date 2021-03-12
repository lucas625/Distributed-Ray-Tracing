FROM python:3.9.0-alpine3.12

# Install dependencies only
COPY requirements.txt /image_generator/
RUN pip install -r /image_generator/requirements.txt

# Setup the rest of the project to be executed
COPY . /image_generator
WORKDIR /image_generator

# Env vars
ENV SECRET_KEY ""
ENV DEBUG false

ENTRYPOINT ["./entrypoint.sh"]
