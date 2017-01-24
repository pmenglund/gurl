[![Build Status](https://travis-ci.org/pmenglund/gurl.svg?branch=master)](https://travis-ci.org/pmenglund/gurl)
[![Coverage Status](https://coveralls.io/repos/github/pmenglund/gurl/badge.svg?branch=master)](https://coveralls.io/github/pmenglund/gurl?branch=master)

# GURL
A simple replacement for `curl` when you want to use it as a docker [`HEALTHCHECK`](https://docs.docker.com/engine/reference/builder/#/healthcheck) instruction.

## Usage

In your `Dockerfile` add a `HEALTHCHECK` line with

    HEALTHCHECK --interval=5m --timeout=3s CMD /gurl http://localhost:8080/

Now you can use `docker ps` to see the health of the container

    docker ps
    CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                   PORTS                    NAMES
    e1b2d3559cc3        gurl                "/server"                2 minutes ago       Up 2 minutes (healthy)                            romantic_leakey

If you want to check the output from the health check, use:

    docker inspect --format='{{json .State.Health}}' <container name> | jq
    {
      "Status": "healthy",
      "FailingStreak": 0,
      "Log": [
        {
          "Start": "2017-01-24T03:31:38.301163386Z",
          "End": "2017-01-24T03:31:38.383919716Z",
          "ExitCode": 0,
          "Output": "alive\n"
        },
        {
          "Start": "2017-01-24T03:31:53.385446984Z",
          "End": "2017-01-24T03:31:53.469891391Z",
          "ExitCode": 0,
          "Output": "alive\n"
        }
      ]
    }
