[![Build Status](https://travis-ci.org/pmenglund/gurl.svg?branch=master)](https://travis-ci.org/pmenglund/gurl)
[![Coverage Status](https://coveralls.io/repos/github/pmenglund/gurl/badge.svg?branch=master)](https://coveralls.io/github/pmenglund/gurl?branch=master)

# GURL
A simple replacement for `curl` when you want to use it as a docker [`HEALTHCHECK`](https://docs.docker.com/engine/reference/builder/#/healthcheck) instruction.

## Usage

In your `Dockerfile` add a `HEALTHCHECK` line with

    HEALTHCHECK --interval=5m --timeout=3s CMD gurl http://localhost:8080/

If you want to check the output

    docker inspect --format='{{json .State.Health}}' <container name> | jq
