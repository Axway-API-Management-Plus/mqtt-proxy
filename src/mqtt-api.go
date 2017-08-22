package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/eclipse/paho.mqtt.golang/packets"
	log "github.com/sirupsen/logrus"
)

func apiPublishMqttPrepare() {
	http.HandleFunc("/topic/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("api: incoming request", r.Header)
		user, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized.", 401)
			return
		}
		var pConnect packets.ConnectPacket
		pConnect.Username = user
		pConnect.Password = []byte(pass)
		pConnect.ClientIdentifier = "api-"
		pConnect.CleanSession = true
		pConnect.ProtocolName = "MQTT"
		pConnect.ProtocolVersion = 5
		session := NewSession()
		err := session.HandleConnect(">", &pConnect)
		if err != nil {
			http.Error(w, "Unauthorized.", 401)
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request.", 400)
			return
		}

		var pPublish packets.PublishPacket
		pPublish.Qos = 1
		pPublish.TopicName = strings.TrimPrefix(r.URL.RequestURI(), "/topic/")
		pPublish.Payload = body

		err = session.HandlePublish(">", &pPublish)
		if err != nil {
			http.Error(w, "Bad Request.", 400)
			return
		}

		broker := "tcp://" + mqttBrokerHost + ":" + strconv.Itoa(mqttBrokerPort)

		/*if strings.HasPrefix(broker, "mqtt://") {
			broker = "tcp://" + strings.TrimPrefix(broker, "mqtt://")
		}*/

		opts := MQTT.NewClientOptions().AddBroker(broker)
		opts.SetClientID(pConnect.ClientIdentifier)
		opts.SetUsername(pConnect.Username)
		opts.SetPassword(string(pConnect.Password))

		//create and start a client using the above ClientOptions
		client := MQTT.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Errorln("Connect", broker, "as", pConnect.ClientIdentifier, token.Error())
			http.Error(w, "Internal Error.", 500)
			log.Errorln(token.Error())
			return
		}

		//subscribe to the topic /go-mqtt/sample and request messages to be delivered
		//at a maximum qos of zero, wait for the receipt to confirm the subscription
		if token := client.Publish(pPublish.TopicName, 0, false, pPublish.Payload); token.Wait() && token.Error() != nil {
			log.Errorln("Publish", pConnect.ClientIdentifier, pPublish.TopicName, token.Error())
			client.Disconnect(0)
			client = nil
			http.Error(w, "Internal Error.", 500)
			log.Errorln(token.Error())
			return
		}

		http.Error(w, "OK.", 200)

		client.Disconnect(0)

	})
}
