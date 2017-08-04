# MQTT-Proxy

## Start
Command Line
```
Usage of ./mqtt-proxy:
  -auth-ca-file string
    	PEM encoded CA s certificate file for the authz/authn service
  -auth-url string
    	URL to the authz/authn service
  -http-enable
    	Enable http protocol (default true)
  -http-host string
    	Listen http port (for http and websockets) (default "0.0.0.0")
  -http-port int
    	Listen http port (for http and websockets) (default 8080)
  -https-cert string
    	Certificate used for https. (default "certs/server.pem")
  -https-enable
    	Enable https protocol (default true)
  -https-host string
    	Listen https port (for https and websockets tls) (default "0.0.0.0")
  -https-key string
    	Key used for https. (default "certs/server.key")
  -https-port int
    	Listen https port (for https and websockets tls) (default 8081)
  -mqtt-broker-host string
    	Host the mqtt server. (default "0.0.0.0")
  -mqtt-broker-password string
    	Password the mqtt server.
  -mqtt-broker-port int
    	Port of the mqtt server (default 1883)
  -mqtt-broker-username string
    	Username of the mqtt server. Reuse incoming one if empty
  -mqtt-enable
    	Enable mqtt protocol (default true)
  -mqtt-host string
    	Mqtt interface to listen to. (default "0.0.0.0")
  -mqtt-port int
    	Mqtt port to listen to. (default 1883)
  -mqtts-cert string
    	Certificate used for mqtt TLS. (default "certs/server.pem")
  -mqtts-enable
    	Enable mqtts protocol (default true)
  -mqtts-host string
    	Mqtts interface to listen to. (default "0.0.0.0")
  -mqtts-key string
    	Key used for mqtt TLS. (default "certs/server.key")
  -mqtts-port int
    	Mqtts port to listen to. (default 1884)
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
- 0.0.4
  - Add HTTPS and MQTTS server support `--mqtts-*` `--https-*`
  - Add HTTPS client support for authz/authn with added cert verification  
  - Remove the 2 steps build: use docker 17.05 build capability

- 0.0.3
  - Add websocket support `--http-host` `--http-port`
  - rename `--mqtt-*` variables to `--mqtt-broker-*`


## Todo
- Broker routing : Configurable Route to the mqtt server on CONNECT
- TLS support broker
- TLS : more TLS options (algo, ....)
- HTTP : Post support
- Cache the publish/receive/subscribe policy check


## Contributing

Please read [Contributing.md](https://github.com/Axway-API-Management-Plus/Common/blob/master/Contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Team

![alt text][Axwaylogo] Axway Team

[Axwaylogo]: https://github.com/Axway-API-Management/Common/blob/master/img/AxwayLogoSmall.png  "Axway logo"


## License
[Apache License 2.0](/LICENSE)
