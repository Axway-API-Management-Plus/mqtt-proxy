# MQTT-Proxy

## Start
Command Line
```
mqtt-proxy --auth-url <AUTH_URL:api-gateway> --mqtt-username <MQTT_USERNAME:guest> --mqtt-password <MQTT_PASSWORD:guest> --mqtt-host <MQTT_HOST> --mqtt-port <MQTT_PORT>
```

Test full environment with docker compose
```
docker-compose -f docker-compose.test.yml up
```

## Configure
When AUTH_URL is set, every mqtt packet (CONNECT, SUBSCRIBE, PUBLISH in/out) are check against a MQTT-AUTH-API :
### `POST $AUTH_URL/connect`
#### Request
```json
{ "Uuid": "",
  "Username" : "",
  "Password" : "",
  "ClientID" : "",
}
```
#### Response 200
```json
{
  "Username" : "Override",
  "Password" : "Override",
  "ClientID" : "Override",
}
```
#### Response Error
The mqtt connection is aborted...

### `POST $AUTH_URL/subscribe`
#### Request
```json
 { "Uuid": "",
   "Username" : "",
   "Password" : "",
   "ClientID" : "",
   "Topic" : ""
}
```
#### Response
```json
{
   "Topic" : "Override"
}
```
#### Response Error
The subscription is cleanly rejected

### `$AUTH_URL/publish`, `$AUTH_URL/receive`
#### Request
```json
{ "Uuid": "",
   "Username" : "",
   "Password" : "",
   "ClientID" : "",
   "Topic" : "",
   "Message": ""
}
```
#### Response
```json
{
   "Topic" : "Override",
   "Message": "Override"
}
```
#### Response Error
The connetion is is aborted (No MQTT Protocol way)

## Build
### Binary:
Prerequisites : `golang`
```sh
make install-deps
make
```

### Docker

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
