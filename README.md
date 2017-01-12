# gurl

A simple replacement for `curl` when you want to use it as a docker `HEALTHCHECK`

## Usage

In your `Dockerfile` add a `HEALTHCHECK` line with

    HEALTHCHECK CMD gurl http://localhost:8080
