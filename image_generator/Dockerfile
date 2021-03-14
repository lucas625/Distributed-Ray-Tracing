FROM python:3.9-slim

# Install dependencies only
RUN pip install --upgrade pip
COPY requirements.txt /image_generator/
RUN pip install -r /image_generator/requirements.txt

# Setup the rest of the project to be executed
COPY . /image_generator
WORKDIR /image_generator

# Env vars
ENV SECRET_KEY ""
ENV DEBUG false

ENTRYPOINT ["./entrypoint.sh"]
