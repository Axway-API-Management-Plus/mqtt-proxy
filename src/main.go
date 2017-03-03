package main

import (
	"net"
	"os"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/namsral/flag"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var host string
var port int
var mqttHost string
var mqttPort int
var mqttUsername string
var mqttPassword string
var authURL string
var Version string
var Build string
var Date string

func main() {
	formatter := new(prefixed.TextFormatter)
	formatter.DisableTimestamp = false
	formatter.ShortTimestamp = false
	formatter.TimestampFormat = "2006-01-02 15:04:05.000000000"
	log.SetFormatter(formatter)
	log.SetLevel(log.DebugLevel)

	flag.IntVar(&port, "port", 1883, "Specify the port to listen to.")
	flag.StringVar(&host, "host", "0.0.0.0", "Specify the interface to listen to.")
	flag.IntVar(&mqttPort, "mqtt-port", 1883, "Specify the port of the mqtt server")
	flag.StringVar(&mqttHost, "mqtt-host", "0.0.0.0", "Specify the host the mqtt server.")
	flag.StringVar(&mqttUsername, "mqtt-username", "", "Specify the username of the mqtt server. Reuse incoming one if empty")
	flag.StringVar(&mqttPassword, "mqtt-password", "", "Specify the password the mqtt server.")
	flag.StringVar(&authURL, "auth-url", "", "Specify the url of the registry service.")
	flag.Parse()

	log.Println("Starting mqtt-proxy @ ", Version, Build, Date)
	log.Println("mqtt server ", mqttHost, mqttPort, mqttUsername, mqttPassword)

	if authURL != "" {
		log.Println("auth connect   : ", authURL+"/connect")
		log.Println("auth publish   : ", authURL+"/publish")
		log.Println("auth subscribe : ", authURL+"/subscribe")
	} else {
		log.Println("auth : no auth url configured : bypassing!")
	}
	// Listen for incoming connections.
	addr := host + ":" + strconv.Itoa(port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Println("Listening on " + addr)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		session := NewSession()
		go session.Stream(conn)
	}
}
