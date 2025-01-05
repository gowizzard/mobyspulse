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
# This is a custom exporter for Docker metrics. Your system running Docker version "27.4.0" and API version "1.47" is being monitored.
container_restart_count{id="d21881ca074683ed47467bd952b4a1c008c59d0d70ef9686641c9cff0257a733",name="mobyspulse",image="ghcr.io/gowizzard/mobyspulse:latest",status="running",health="healthy",created="1736067482",started_at="1736067482"} 0
```