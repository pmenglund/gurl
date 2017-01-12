# GURL
[![Build Status](https://travis-ci.org/pmenglund/gurl.svg?branch=master)](https://travis-ci.org/pmenglund/gurl)

A simple replacement for `curl` when you want to use it as a docker [`HEALTHCHECK`](https://docs.docker.com/engine/reference/builder/#/healthcheck) instruction.

## Usage

In your `Dockerfile` add a `HEALTHCHECK` line with

    HEALTHCHECK --interval=5m --timeout=3s CMD gurl http://localhost:8080/

If you want to check the output

    docker inspect <container>
