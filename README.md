# MQTT Proxy
![alt text](https://img.shields.io/docker/automated/davinci1976/mqtt-proxy.svg)
![alt text](https://img.shields.io/docker/build/davinci1976/mqtt-proxy.svg)

Proxy to apply Axway API Gateway policies (authN, authZ, content manipulation,...) on MQTT protocol for any MQTT broker

## API Management Version Compatibilty
This artefact can be used with every API Management Plus version

## Prerequisites
- docker 17.06 (and docker-compose)
- MQTT Broker   : activemq / rabbitmq / mosquitto / ...
- Policy Engine : Axway API Gateway or custom engine (see [./tests/policy] )

## Configure your policy engine
When AUTH_URL is set, every mqtt packet (CONNECT, SUBSCRIBE, PUBLISH in/out) are check against a [AUTH_API](./AUTH_API.md)

See Axway API Gateway samples for mqtt in : `./api-gateway-policies/mqtt-proxy-apigw-policy.xml`
In API Gateway Policy Studio, please use Import Configuration Fragment to upload the policy

## Command-line / Environment options
```sh
docker run -it --rm mqtt-proxy davinci1976/mqtt-proxy mqtt-proxy --help
```

## Quickstart
### Standalone (mqtt-proxy only)
```sh
   docker run -it --rm davinci1976/mqtt-proxy mqtt-proxy --auth-url http://apigtw:8065/mqtt --mqtt-broker-host my-mqtt-broker
```

### Full environment (mqtt-proxy + broker + custom policy engine)
```sh
docker-compose -f docker-compose.yml up
```

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
-or-
```
make docker
```

## Test
```
make docker-test
```

## Changelog
- 0.0.5
  - Add HTTP(s) API for MQTT Publish `/api/topics/:topic` with basic authentication
- 0.0.4
  - Add HTTPS and MQTTS server support `--mqtts-*` `--https-*`
  - Add HTTPS client support for authz/authn with added cert verification  
  - Remove the 2 steps build: use docker 17.05 build capability
- 0.0.3
  - Add websocket support `--http-host` `--http-port`
  - rename `--mqtt-*` variables to `--mqtt-broker-*`


## Limitations/Caveats
- No broker routing : Only one broker per mqtt-proxy instance
  - no configurable routes
- No TLS support for broker
- No additional TLS options supported between the client and mqtt-proxy (algo, ....)
- No cache for publish/receive/subscribe policy check

## Contributing

Please read [Contributing.md](https://github.com/Axway-API-Management-Plus/Common/blob/master/Contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Team

![alt text][Axwaylogo] Axway Team

[Axwaylogo]: https://github.com/Axway-API-Management/Common/blob/master/img/AxwayLogoSmall.png  "Axway logo"


## License
[Apache License 2.0](/LICENSE)
