# MQTT-Proxy

## Start
Command Line
```sh
mqtt-proxy [--auth-url <AUTH_URL:api-gateway>] \
           [--mqtt-broker-username <MQTT_BROKER_USERNAME:guest>] \
           [--mqtt-broker-password <MQTT_BROKER_PASSWORD:guest>] \
           [--mqtt-broker-host <MQTT_BROKER_HOST:localhost>] \
           [--mqtt-broker-port <MQTT_BROKER_PORT:1883>] \
           \
           [--host <HOST:0.0.0.0>] \
           [--port <PORT:1883>] \
           [--http-host <HTTP_HOST:0.0.0.0>] \
           [--http-port <HTTP_PORT:8080>]
```

Test full environment with docker
```
docker-compose -f docker-compose.yml up
```

## Configure

When AUTH_URL is set, every mqtt packet (CONNECT, SUBSCRIBE, PUBLISH in/out) are check against a [AUTH_API](./AUTH_API.md)


## Build standalone binary:
Prerequisites : `golang`
```sh
make install-deps
make
```

## Build docker image

```
docker build -t mqtt-proxy:dev .
```

```
make docker
  -or-
docker build -t mqtt-proxy:dev .
docker run --rm mqtt-proxy:dev tar cz mqtt-proxy >mqtt-proxy.tar.gz
docker build -t mqtt-proxy -f Dockerfile.small .
```

## Test

```
make docker-test
```

## Changelog
* 0.0.3
  * Add websocket support `--http-host` `--http-port`
  * rename `--mqtt-*` variables to `--mqtt-broker-*`

## Todo
- Configurable Route to the mqtt server on CONNECT
- TLS support
- Websocket support
- HTTP Post support
- Cache the publish/subscribe policy check
