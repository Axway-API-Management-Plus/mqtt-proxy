package main

import (
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/eclipse/paho.mqtt.golang/packets"
)

func (session *Session) ForwardMQTTPacket(r net.Conn, w net.Conn) error {
	cp, err := packets.ReadPacket(r)
	if err != nil {
		if !session.closed {
			log.Errorln("Session", session.id, "- Error reading MQTT packet", err)
		}
		return err
	}
	log.Println("Session", session.id, "- Forward MQTT packet", cp.String())

	if authURL != "" {
		switch p := cp.(type) {
		case *packets.ConnectPacket: /*Outbound only*/
			err = session.HandleConnect(p, r, w)
		case *packets.SubscribePacket: /*Outbound only*/
			err = session.HandleSubscribe(p, r, w)
		case *packets.PublishPacket: /*Inbound/Outbound only*/
			err = session.HandlePublish(p, r, w)
		default:
			err = nil
		}
	} else {
		err = nil
	}

	if err != nil {
		log.Println("Session", session.id, "- Forward MQTT packet", err)
		return err
	}

	err = cp.Write(w)
	if err != nil {
		log.Errorln("Session", session.id, "- Error writing MQTT packet", err)
		return err
	}
	return nil
}
