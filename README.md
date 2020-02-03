# docker-applications

Run docker-compose commands on multiple apps.

[![CircleCI](https://circleci.com/gh/Lajule/docker-applications/tree/master.svg?style=svg&circle-token=3bb0c1914c37e942e3b5597f4789cac8943c67e2)](https://circleci.com/gh/Lajule/docker-applications/tree/master)

Don't write this anymore :

```sh
docker-compose -f $HOME/front/docker-compose.yml -f $HOME/api/api.yml -f /varnish/docker-compose.yml up -d --build
```

Instead write :

```sh
docker-applications front up -d --build
```

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
