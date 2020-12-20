# Distributed-Ray-Tracing

[![Build Status](https://travis-ci.org/lucas625/Distributed-Ray-Tracing.svg?branch=master)](https://travis-ci.org/lucas625/Distributed-Ray-Tracing) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/c97b506fef9f4eb8a23da10b04a04fb1)](https://www.codacy.com/manual/lucas625/Distributed-Ray-Tracing?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=lucas625/Distributed-Ray-Tracing&amp;utm_campaign=Badge_Grade)[![codecov](https://codecov.io/gh/lucas625/Distributed-Ray-Tracing/branch/master/graph/badge.svg)](https://codecov.io/gh/lucas625/Distributed-Ray-Tracing)

The implementation of a distributed ray tracing based on microservices as [Lucas Aurelio](https://github.com/lucas625) is completion of course work.

## Table of Contents

[[_TOC_]]

## Team

Developer: [Lucas Aurelio](https://github.com/lucas625)

## Requirements

- [Docker](https://docs.docker.com/desktop/)

## Guidelines

- [The Twelve-Factor App](https://12factor.net/)

## Installation

Follow the next steps to install the dependencies.

### Ray-Tracing

To set ray tracing project use:

```sh
# Build the image
docker build -t drt-ray-tracing .

# Run the container
docker run -p 8081:8081 --name drt-ray-tracing-container -v $(pwd):/ray-tracing --rm drt-ray-tracing
```

## Run the project

After completing the installation of the project, follow the next steps to execute it.

## Testing

Follow the next steps to test the application.

### Ray-Tracing Test

Inside **drt-ray-tracing-container** run:

```sh
go test ./...
```
