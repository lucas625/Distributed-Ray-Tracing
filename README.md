# Distributed-Ray-Tracing

[![Build Status](https://travis-ci.org/lucas625/Distributed-Ray-Tracing.svg?branch=master)](https://travis-ci.org/lucas625/Distributed-Ray-Tracing) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c97b506fef9f4eb8a23da10b04a04fb1)](https://www.codacy.com/manual/lucas625/Distributed-Ray-Tracing?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=lucas625/Distributed-Ray-Tracing&amp;utm_campaign=Badge_Grade)[![codecov](https://codecov.io/gh/lucas625/Distributed-Ray-Tracing/branch/master/graph/badge.svg)](https://codecov.io/gh/lucas625/Distributed-Ray-Tracing)

The implementation of a distributed ray tracing based on microservices as [Lucas Aurelio](https://github.com/lucas625) is completion of course work.

## Table of Contents

- [Distributed-Ray-Tracing](#distributed-ray-tracing)
  - [Table of Contents](#table-of-contents)
  - [Team](#team)
  - [Requirements](#requirements)
  - [Guidelines](#guidelines)
  - [Deploy](#deploy)
  - [Testing](#testing)
    - [Ray-Tracing Test](#ray-tracing-test)
  - [Examples](#examples)

## Team

Developer: [Lucas Aurelio](https://github.com/lucas625)

## Requirements

- [Docker](https://docs.docker.com/desktop/)

## Guidelines

- [The Twelve-Factor App](https://12factor.net/)

## Deploy

Follow the steps on [deploy file](deploy.md).

## Testing

Follow the next steps to test the application.

### Ray-Tracing Test

- Automatic tests
Inside **drt-ray-tracing-container** run:

```sh
go test ./...
```

- Manual Tests

```bash
# Build the environment variables
# Remember to change the values as needed.
./build_docker_env_vars.sh

# Run the services
./run_docker.sh
```

## Examples

You can see samples for the path tracing on the `sample_objects` folder. The structure of the `.JSON` files is the same as the beans is structure.
Please keep in mind that if you want to access the API directly, you will need to provide more information than just the data contained in the sample objects. This data is the `pixelScreen` and the `pathTracingParameters`. You can see how to build this data in `frontend/src/views/RayTracingView.vue`.
