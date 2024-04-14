<div align="center">


<img src="https://github.com/gowizzard/mobyspulse/assets/30717818/4e6286f2-eec3-416c-ade4-1819f331ae49" alt="Moby's Pulse" width="250" style="margin:50px 0;">

# Moby's Pulse

[![Go Test](https://github.com/gowizzard/mobyspulse/actions/workflows/go-test.yml/badge.svg)](https://github.com/gowizzard/mobyspulse/actions/workflows/go-test.yml) [![Docker Test](https://github.com/gowizzard/mobyspulse/actions/workflows/docker-test.yml/badge.svg)](https://github.com/gowizzard/mobyspulse/actions/workflows/docker-test.yml) [![Docker Build](https://github.com/gowizzard/mobyspulse/actions/workflows/docker-build.yml/badge.svg)](https://github.com/gowizzard/mobyspulse/actions/workflows/docker-build.yml) [![Pull Request Labels](https://github.com/gowizzard/mobyspulse/actions/workflows/pull-request-labels.yml/badge.svg)](https://github.com/gowizzard/mobyspulse/actions/workflows/pull-request-labels.yml)

A little prometheus exporter to get the restarts of containers, with multiple attributes to get specific container information.

</div>

## Installation

The easiest way to install Moby's Pulse is to use the provided Docker image. You can pull the image from the GitHub Container Registry:

```bash
docker pull ghcr.io/gowizzard/mobyspulse:latest
```

## Usage

To start the exporter, you can use the following docker compose file:

```yaml
services:
    mobyspulse:
        container_name: mobyspulse
        environment:
            - BASIC_AUTH_USERNAME=moby
            - BASIC_AUTH_PASSWORD=pulse
        volumes:
            - /var/run/docker.sock:/var/run/docker.sock
        ports:
            - "3000:3000"
        image: ghcr.io/gowizzard/mobyspulse:latest
```

If you want to use the exporter without basic auth, you can remove the environment variables. The exporter will be available on port 3000.

### Metrics

The exporter provides the following metrics:

```text
# Moby's Pulse - Docker metrics exporter for Prometheus.
# This is a custom exporter for Docker metrics. Your system running Docker version "26.0.1" and API version "1.45" is being monitored.
container_restart_count{id="2c35a3500c6384cd88b1cb30182ba39e54c74e4047515d2e6db662aaa7116bb2",name="mobyspulse",image="mobyspulse-mobyspulse:latest",created="1713079803",started_at="1713079803"} 0
```