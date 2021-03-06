# Distributed-Ray-Tracing

[![Build Status](https://travis-ci.org/lucas625/Distributed-Ray-Tracing.svg?branch=master)](https://travis-ci.org/lucas625/Distributed-Ray-Tracing) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c97b506fef9f4eb8a23da10b04a04fb1)](https://www.codacy.com/manual/lucas625/Distributed-Ray-Tracing?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=lucas625/Distributed-Ray-Tracing&amp;utm_campaign=Badge_Grade)[![codecov](https://codecov.io/gh/lucas625/Distributed-Ray-Tracing/branch/master/graph/badge.svg)](https://codecov.io/gh/lucas625/Distributed-Ray-Tracing)

The implementation of a distributed ray tracing based on microservices as [Lucas Aurelio](https://github.com/lucas625) is completion of course work.

## Table of Contents

- [Distributed-Ray-Tracing](#distributed-ray-tracing)
  - [Table of Contents](#table-of-contents)
  - [Team](#team)
  - [Requirements](#requirements)
  - [Guidelines](#guidelines)
  - [Run the project](#run-the-project)
    - [Building Images](#building-images)
    - [Run the images](#run-the-images)
  - [Testing](#testing)
    - [Ray-Tracing Test](#ray-tracing-test)
  - [Examples](#examples)

## Team

Developer: [Lucas Aurelio](https://github.com/lucas625)

## Requirements

- [Docker](https://docs.docker.com/desktop/)

## Guidelines

- [The Twelve-Factor App](https://12factor.net/)

## Run the project

After completing the installation of the project, follow the next steps to execute it.

### Building Images

```sh
# Build the ray tracing image
docker build -t drt-ray-tracing -f ray-tracing/Dockerfile ray-tracing
```

### Run the images

```sh
# Run the ray tracing image
docker run -p 8081:8081 --rm --name drt-ray-tracing-container drt-ray-tracing
```

## Testing

Follow the next steps to test the application.

### Ray-Tracing Test

Inside **drt-ray-tracing-container** run:

```sh
go test ./...
```

## Examples

You can see samples for the path tracing on the `sample_objects` folder. The structure of the `.JSON` files is the same as the beans is structure.
