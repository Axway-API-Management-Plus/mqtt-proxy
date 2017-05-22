#

## Setup
- Pane 1 (bottom right): see mqttproxy logs
```
git clone git@git.ecd.axway.int:jdavanne/mqtt-proxy.git
cd mqtt-proxy
docker-compose up
docker logs -f mqttproxy_mqtt-proxy_1
```
- Pane 2 (upper left) : subscriber (Thing)
mqttcli sub -p 1883 --host localhost --user test1 --password goodpass -q 1 -t test -d

mqttcli pub -p 1883 --host localhost --user test2 --password goodpass -q 1 -t test -m "Hello"

```

1. Control on user/password
- `mqttcli sub -p 1883 --host localhost --user test1 --password goodpass -q 1 -t test -d`
2. Control on subscribed topic
- `mqttcli sub -p 1883 --host localhost --user test1 --password goodpass -q 1 -t bad_topic_on_subscribe -d`
3. Control on message sent
- `mqttcli pub -p 1883 --host localhost --user test1 --password goodpass -t test -m "bad_payload_on_publish"`
4. Control on message received
5. Modify topics
- `mqttcli pub -p 1883 --host localhost --user test1 --password goodpass -t "rename_topic_on_publish" -m "message2"`
6. Modify messages
- on publish
  - `mqttcli pub -p 1883 --host localhost --user test1 --password goodpass -t test -m "alter_message_on_publish"`
- on receive
- ` mqttcli pub -p 1883 --host localhost --user test2 --password goodpass -t "rename_topic_on_receive" -m "message2"
