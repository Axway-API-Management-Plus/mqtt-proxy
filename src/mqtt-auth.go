package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/eclipse/paho.mqtt.golang/packets"
)

type MQTTConnect struct {
	Uuid             string
	Username         string
	Password         string
	ClientIdentifier string
	CleanSession     bool
	ProtocolName     string
	ProtocolVersion  int
}

type MQTTConnectResponse struct {
	Username         string
	Password         string
	ClientIdentifier string
}

type MQTTSubscribe struct {
	Uuid             string
	Username         string
	ClientIdentifier string
	Topic            string
	Qos              int
}

type MQTTSubscribeResponse struct {
	Topic string
}

type MQTTPublish struct {
	Uuid             string
	Username         string
	ClientIdentifier string
	Topic            string
	Qos              int
	Payload          string
}

type MQTTPublishResponse struct {
	Topic   string
	Payload string
}

func request(uri string, request interface{}, response interface{}) (int, error) {
	//req.Header.Add("API-KEY", "tenant:admin")
	jData, err := json.Marshal(request)
	if err != nil {
		return 0, err
	}

	tr := &http.Transport{
	//TLSClientConfig:    &tls.Config{RootCAs: pool},
	//DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", authURL+uri, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("API-KEY", "tenant:admin")
	//req.Header.Add("Authorization", "")
	req.Body = ioutil.NopCloser(bytes.NewReader(jData))

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	log.Println("auth response", resp.StatusCode, string(body))
	if response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return 0, err
		}
	}

	return resp.StatusCode, nil
}

func (session *Session) HandleConnect(p *packets.ConnectPacket, r net.Conn, w net.Conn) error {
	log.Println("Session", session.id, "- CONNECT")
	var resp MQTTConnectResponse
	rq := MQTTConnect{session.id, p.Username, string(p.Password), p.ClientIdentifier, p.CleanSession, p.ProtocolName, int(p.ProtocolVersion)}
	code, err := request("/connect", rq, &resp)
	if err != nil {
		log.Errorln("Session", session.id, "- Error getting connect authorization", err)
		return err
	}

	if code != 200 {
		return errors.New("Connect Not Authorized")
	}
	if mqttUsername != "" {
		p.Username = mqttUsername
		p.Password = []byte(mqttPassword)
	}

	//Override information
	if resp.ClientIdentifier != "" && resp.ClientIdentifier != p.ClientIdentifier {
		log.Println("Session", session.id, "- CONNECT alter ClientIdentifier", p.ClientIdentifier, "-->", resp.ClientIdentifier)
		p.ClientIdentifier = resp.ClientIdentifier
	}
	if resp.Username != "" && resp.Username != p.Username {
		log.Println("Session", session.id, "- CONNECT alter Username", p.Username, "-->", resp.Username)
		p.Username = resp.Username
	}
	if resp.Password != "" && resp.Password != string(p.Password) {
		log.Println("Session", session.id, "- CONNECT alter Password")
		p.Password = []byte(resp.Password)
	}

	session.Username = p.Username
	session.ClientIdentifier = p.ClientIdentifier

	return nil
}

func (session *Session) HandleSubscribe(p *packets.SubscribePacket, r net.Conn, w net.Conn) error {
	log.Println("Session", session.id, "- SUBSCRIBE", p.Topics, p.Qos)
	var resp MQTTSubscribeResponse
	topics := p.Topics
	for i := range p.Topics {
		rq := MQTTSubscribe{session.id, session.Username, session.ClientIdentifier, p.Topics[i], int(p.Qos)}
		code, err := request("/subscribe", rq, &resp)

		if err != nil {
			log.Errorln("Session", session.id, "- Error getting subscribe authorization", err)
			return err
		}
		if code != 200 {
			cp2 := packets.NewControlPacket(packets.Suback)
			suback := cp2.(*packets.SubackPacket)
			suback.ReturnCodes = []byte{packets.ErrRefusedNotAuthorised}
			err := suback.Write(r)
			if err != nil {
				log.Errorln("Session", session.id, "- Error writing subscribe ack error message", err)
			}
			return errors.New("Subscribe Not Authorized")
		}

		if resp.Topic != "" && resp.Topic != topics[i] {
			log.Println("Session", session.id, "- SUBSCRIBE alter topic", i, topics[i], "-->", resp.Topic)
			topics[i] = resp.Topic
		}
	}

	p.Topics = topics

	return nil
}

func (session *Session) HandlePublish(p *packets.PublishPacket, r net.Conn, w net.Conn) error {
	action := "PUBLISH"
	uri := "/publish"
	if w == session.inbound {
		action = "RECEIVE"
		uri = "/receive"
	}
	log.Println("Session", session.id, "- "+action, r.RemoteAddr().String(), w.RemoteAddr().String())
	log.Println("Session", session.id, "- "+action, p.TopicName, p.Qos, string(p.Payload))
	rq := MQTTPublish{session.id, session.Username, session.ClientIdentifier, p.TopicName, int(p.Qos), string(p.Payload)}
	var resp MQTTPublishResponse
	code, err := request(uri, rq, &resp)

	if err != nil {
		log.Errorln("Session", session.id, "- Error getting Publish authorization", err)
		return err
	}
	if code != 200 {
		return errors.New(action + " Not Authorized")
	}
	if resp.Topic != "" && resp.Topic != p.TopicName {
		log.Println("Session", session.id, "- "+action+" alter topic", p.TopicName, "-->", resp.Topic)
		p.TopicName = resp.Topic
	}
	if resp.Payload != "" && resp.Payload != string(p.Payload) {
		log.Println("Session", session.id, "- "+action+"alter topic", p.Payload, "-->", resp.Payload)
		p.Payload = []byte(resp.Payload)
	}
	return nil
}
