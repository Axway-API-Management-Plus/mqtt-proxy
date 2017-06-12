# MQTT-Proxy

## Start
Command Line
```sh
Usage of ./mqtt-proxy:
  -auth-url string
    	Specify the url of the registry service.
  -host string
    	Specify the mqtt interface to listen to. (default "0.0.0.0")
  -http-host string
    	Listen http port (for http and websockets) (default "0.0.0.0")
  -http-port int
    	Listen http port (for http and websockets) (default 8080)
  -https-cert string
    	Specify the certificate used for https. (default "certs/server.pem")
  -https-host string
    	Listen https port (for https and websockets tls) (default "0.0.0.0")
  -https-key string
    	Specify the key used for https. (default "certs/server.key")
  -https-port int
    	Listen https port (for https and websockets tls) (default 8081)
  -mqtt-broker-host string
    	Specify the host the mqtt server. (default "0.0.0.0")
  -mqtt-broker-password string
    	Specify the password the mqtt server.
  -mqtt-broker-port int
    	Specify the port of the mqtt server (default 1883)
  -mqtt-broker-username string
    	Specify the username of the mqtt server. Reuse incoming one if empty
  -mqtts-cert string
    	Specify the certificate used for mqtt TLS. (default "certs/server.pem")
  -mqtts-host string
    	Specify the mqtts interface to listen to. (default "0.0.0.0")
  -mqtts-key string
    	Specify the key used for mqtt TLS. (default "certs/server.key")
  -mqtts-port int
    	Specify the mqtts port to listen to. (default 1884)
  -port int
    	Specify the mqtt port to listen to. (default 1883)
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
Prerequisites : `docker 17.05`

```
docker build -t mqtt-proxy .
```

```
make docker
```

## Test

```
make docker-test
```

## Changelog
* 0.0.3
  * Add websocket support `--http-host` `--http-port`
  * rename `--mqtt-*` variables to `--mqtt-broker-*`
* 0.0.4
  * Add TLS support for mqtt and https listeners `--mqtts-*` `--https-*``
  * Remove the 2 steps build: use docker 17.05 build capability

## Todo
- Broker routing : Configurable Route to the mqtt server on CONNECT
- TLS support broker
- TLS : more TLS options (algo, ....)
- HTTP : Post support
- Cache the publish/receive/subscribe policy check

