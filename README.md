# docker-applications

Run docker-compose commands on multiple apps.

## Usage

Type the following command `docker-applications -h` to display this help message:

```
docker-applications - Run docker-compose commands on multiple apps

Usage:
  docker-applications [flags]

Flags:
  -f, --file string   config file (default is ./docker-applications.yml)
  -h, --help          help for docker-applications
```

## Configuration

This is an example of a possible configuration:

```yaml
version: '1'
applications:
  front:
    dir: ${HOME}/front
    depends_on:
      - varnish
      - api

  api:
    dir: ${HOME}/api
    file: api.yml
    depends_on:
      - api

  varnish:
    dir: /varnish
```
