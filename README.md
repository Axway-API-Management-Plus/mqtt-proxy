# MQTT-Proxy

## Start
Command Line
```
mqtt-proxy --auth-url <AUTH_URL:api-gateway> --mqtt-username <MQTT_USERNAME:guest> --mqtt-password <MQTT_PASSWORD:guest> --mqtt-host <MQTT_HOST> --mqtt-port <MQTT_PORT>
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

## Todo
- Configurable Route to the mqtt server on CONNECT
- TLS support
- Websocket support
- HTTP Post support
- Cache the publish/subscribe policy check
