package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/namsral/flag"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var host string
var port int

var mqttsHost string
var mqttsPort int
var mqttsCert string
var mqttsKey string

var httpHost string
var httpPort int
var httpsHost string
var httpsPort int
var httpsCert string
var httpsKey string

var mqttBrokerHost string
var mqttBrokerPort int
var mqttBrokerUsername string
var mqttBrokerPassword string
var authURL string
var Version string
var Build string
var Date string

func main() {
	formatter := new(prefixed.TextFormatter)
	formatter.DisableTimestamp = false
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000000000"
	log.SetFormatter(formatter)
	log.SetLevel(log.DebugLevel)

	flag.IntVar(&port, "port", 1883, "Specify the mqtt port to listen to.")
	flag.StringVar(&host, "host", "0.0.0.0", "Specify the mqtt interface to listen to.")
	flag.IntVar(&mqttsPort, "mqtts-port", 1884, "Specify the mqtts port to listen to.")
	flag.StringVar(&mqttsHost, "mqtts-host", "0.0.0.0", "Specify the mqtts interface to listen to.")
	flag.StringVar(&mqttsCert, "mqtts-cert", "certs/server.pem", "Specify the certificate used for mqtt TLS.")
	flag.StringVar(&mqttsKey, "mqtts-key", "certs/server.key", "Specify the key used for mqtt TLS.")

	flag.StringVar(&httpHost, "http-host", "0.0.0.0", "Listen http port (for http and websockets)")
	flag.IntVar(&httpPort, "http-port", 8080, "Listen http port (for http and websockets)")
	flag.StringVar(&httpsHost, "https-host", "0.0.0.0", "Listen https port (for https and websockets tls)")
	flag.IntVar(&httpsPort, "https-port", 8081, "Listen https port (for https and websockets tls)")
	flag.StringVar(&httpsCert, "https-cert", "certs/server.pem", "Specify the certificate used for https.")
	flag.StringVar(&httpsKey, "https-key", "certs/server.key", "Specify the key used for https.")

	flag.IntVar(&mqttBrokerPort, "mqtt-broker-port", 1883, "Specify the port of the mqtt server")
	flag.StringVar(&mqttBrokerHost, "mqtt-broker-host", "0.0.0.0", "Specify the host the mqtt server.")
	flag.StringVar(&mqttBrokerUsername, "mqtt-broker-username", "", "Specify the username of the mqtt server. Reuse incoming one if empty")
	flag.StringVar(&mqttBrokerPassword, "mqtt-broker-password", "", "Specify the password the mqtt server.")

	flag.StringVar(&authURL, "auth-url", "", "Specify the url of the registry service.")
	flag.Parse()

	log.Println("Starting mqtt-proxy @ ", Version, Build, Date)
	log.Println("mqtt server ", mqttBrokerHost, mqttBrokerPort, mqttBrokerUsername, mqttBrokerPassword)

	if authURL != "" {
		log.Println("auth connect   : ", authURL+"/connect")
		log.Println("auth publish   : ", authURL+"/publish")
		log.Println("auth subscribe : ", authURL+"/subscribe")
	} else {
		log.Println("auth : no auth url configured : bypassing!")
	}

	wsMqttPrepare()
	go wsMqttListen()
	go wssMqttListen()

	go mqttListen()
	mqttsListen()
}
