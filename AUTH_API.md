# MQTT AUTH API

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
The connection is is aborted (No MQTT Protocol way)
