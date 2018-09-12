# Description

[![Build Status](https://travis-ci.org/Axway-API-Management-Plus/mqtt-proxy.svg?branch=master)](https://travis-ci.org/Axway-API-Management-Plus/mqtt-proxy)

The MQTT-Proxy itself sits between the MQTT-Provider & -Consumer and intercepts incoming MQTT-Commands, with the ability to call a REST-API at the API-Gateway. With that, it is for instance possible to validate, that a certain MQTT-Consumer can subscribe to a topic, as the API-Gateway can easily validate the Subscription-Request using a database, another downstream API, whatever.


To get an overview you can also watch our MQTT video series:  
[![MQTT Proxy and Trigger videos](https://img.youtube.com/vi/8RoElGdBVxY/1.jpg)](https://www.youtube.com/playlist?list=PLSlCpG9zsECrWZLocBzr3MM8AAatkArUF)  

![alt text][Image1]

[Image1]: https://github.com/Axway-API-Management-Plus/mqtt-proxy/blob/master/readme/mqtt-proxy01.png "Image1"


The diagram below details more flows between the various components:

![alt text][Image2]

[Image2]: https://github.com/Axway-API-Management-Plus/mqtt-proxy/blob/master/readme/mqtt-proxy02.png "Image2"



## API Management Version Compatibilty
This artefact can be used with every API Management Plus version

## Prerequisites
- docker 17.06 (and docker-compose)
- MQTT Broker   : activemq / rabbitmq / mosquitto / ...
- Policy Engine : Axway API Gateway or custom engine (see [./tests/policy] ) exposing the REST-API to use

## Configure your policy engine
The first important step is to tell the MQTT-Proxy, the location of the REST-API to use. This is configured using the Docker Environment-Variable AUTH_URL and when configured every mqtt packet (CONNECT, SUBSCRIBE, PUBLISH in/out) is checked against this endpoint. The REST-API calls exeuted by the MQTT-Proxy are described here: [AUTH_API](./AUTH_API.md)

We are providing some samples REST-APIs for the Axway API Gateway in: `./api-gateway-policies/mqtt-proxy-apigw-policy.xml`
To use it, just "Import the Configuration Fragment" using Axway API Gateway Policy Studio and deploy this to your API-Gateway.

## Command-line / Environment options
The MQQT-Proxy binary comes with a number of command-line options to enable, disable or control certain features. To review all possible command-line options the MQTT-Proxy provides, please execute the following command:
```sh
docker run -it --rm davinci1976/mqtt-proxy mqtt-proxy --help
```

## Quickstart
To use the MQTT-Proxy you have multiple ways to start it, either using your own MQTT-Broker or an MQTT-Broker included in this asset and started as a Docker-Container.  

### Standalone (mqtt-proxy only)
This means, that no MQTT-Broker is started, hence you have to configure the location of your existing running MQTT-Broker. 
This is a simple example without user-authentication against the MQTT-Broker:
```sh
docker run -it --rm -e AUTH_URL=http://api-host:8080/mqtt -e MQTT_BROKER_HOST=my-mqtt-broker -p 1883:1883 davinci1976/mqtt-proxy
```
An example would be:
```sh
docker run -it --rm -e AUTH_URL=http://172.17.0.1:8080/mqtt -e PORT=1884 -e MQTT_BROKER_HOST=172.10.1 -p 1884:1883 davinci1976/mqtt-proxy
```
or using some of command line options provided by the mqtt-proxy binary:
```sh
docker run -it --rm -e AUTH_URL=http://172.17.0.1:8080/mqtt -e MQTT_BROKER_HOST=172.10.1 -p 1884:1883 davinci1976/mqtt-proxy mqtt-proxy -mqtt-port 1884 -mqtts-port 1885
```
The following environment variables including default values are supported (please the the --help for all command line options):
```
PORT 1883
MQTT_BROKER_HOST 0.0.0.0
MQTT_BROKER_PORT 1883
MQTT_BROKER_USERNAME guest
MQTT_BROKER_PASSWORD guest
AUTH_URL ""
```

### Full environment (mqtt-proxy + broker + custom policy engine)
This brings up a complete environment, including the MQTT-Proxy, MQTT-Broker (RabbitMq, Mosquitto, ActiveMQ) and a Node.js based REST-API listening on http://policy:3000/mqtt for testing purposes:
```sh
docker-compose -f docker-compose.yml up
```
This is much more than needed, but you can adjust the docker-compose.yml file according to your needs.

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
- No HTTP API to publish a MQTT message on a topic `/topics/:topic?qos=:qos`
   (like http://docs.aws.amazon.com/iot/latest/developerguide/protocols.html#http)
- No cache for publish/receive/subscribe policy check

## Contributing

Please read [Contributing.md](https://github.com/Axway-API-Management-Plus/Common/blob/master/Contributing.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Team

![alt text][Axwaylogo] Axway Team

[Axwaylogo]: https://github.com/Axway-API-Management/Common/blob/master/img/AxwayLogoSmall.png  "Axway logo"


## License
[Apache License 2.0](/LICENSE)
