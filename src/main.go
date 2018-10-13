package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var mqttHost string
var mqttPort int
var mqttEnable bool

var mqttsHost string
var mqttsPort int
var mqttsCert string
var mqttsKey string
var mqttsEnable bool

var httpHost string
var httpPort int
var httpEnable bool

var httpsHost string
var httpsPort int
var httpsCert string
var httpsKey string
var httpsEnable bool

var mqttBrokerHost string
var mqttBrokerPort int
var mqttBrokerUsername string
var mqttBrokerPassword string
var authURL string
var authCAFile string
var Version string
var Build string
var Date string

var logLevel string

var authClient *http.Client

func main() {
	formatter := new(prefixed.TextFormatter)
	formatter.DisableTimestamp = false
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000000000"
	log.SetFormatter(formatter)

	flag.String(flag.DefaultConfigFlagname, "", "Path to config file")

	flag.IntVar(&mqttPort, "mqtt-port", 1883, "Mqtt port to listen to.")
	flag.StringVar(&mqttHost, "mqtt-host", "0.0.0.0", "Mqtt interface to listen to.")
	flag.BoolVar(&mqttEnable, "mqtt-enable", true, "Enable mqtt protocol")

	flag.IntVar(&mqttsPort, "mqtts-port", 8883, "Mqtts port to listen to.")
	flag.StringVar(&mqttsHost, "mqtts-host", "0.0.0.0", "Mqtts interface to listen to.")
	flag.StringVar(&mqttsCert, "mqtts-cert", "certs/server.pem", "Certificate used for mqtt TLS.")
	flag.StringVar(&mqttsKey, "mqtts-key", "certs/server.key", "Key used for mqtt TLS.")
	flag.BoolVar(&mqttsEnable, "mqtts-enable", true, "Enable mqtts protocol")

	flag.StringVar(&httpHost, "http-host", "0.0.0.0", "Listen http port (for http and websockets)")
	flag.IntVar(&httpPort, "http-port", 8080, "Listen http port (for http and websockets)")
	flag.BoolVar(&httpEnable, "http-enable", true, "Enable http protocol")

	flag.StringVar(&httpsHost, "https-host", "0.0.0.0", "Listen https port (for https and websockets tls)")
	flag.IntVar(&httpsPort, "https-port", 8081, "Listen https port (for https and websockets tls)")
	flag.StringVar(&httpsCert, "https-cert", "certs/server.pem", "Certificate used for https.")
	flag.StringVar(&httpsKey, "https-key", "certs/server.key", "Key used for https.")
	flag.BoolVar(&httpsEnable, "https-enable", true, "Enable https protocol")

	flag.IntVar(&mqttBrokerPort, "mqtt-broker-port", 1883, "Port of the mqtt server")
	flag.StringVar(&mqttBrokerHost, "mqtt-broker-host", "0.0.0.0", "Host the mqtt server.")
	flag.StringVar(&mqttBrokerUsername, "mqtt-broker-username", "", "Username of the mqtt server. Reuse incoming one if empty")
	flag.StringVar(&mqttBrokerPassword, "mqtt-broker-password", "", "Password the mqtt server.")

	flag.StringVar(&authURL, "auth-url", "", "URL to the authz/authn service")
	flag.StringVar(&authCAFile, "auth-ca-file", "", "PEM encoded CA's certificate file for the authz/authn service")

	flag.StringVar(&logLevel, "log-level", "info", "Log level [debug, info, warn, error](default: info)")
	flag.Parse()


	switch logLevel{
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
		logLevel = "debug"
	}

	log.Println("Starting mqtt-proxy @ ", Version, Build, Date)
	log.Println("mqtt server ", mqttBrokerHost, mqttBrokerPort, mqttBrokerUsername, mqttBrokerPassword)

	log.Println("LogLevel:", logLevel)

	if authURL != "" {
		log.Println("auth url : ", authURL)
		log.Debugln("auth connect   : ", authURL+"/connect")
		log.Debugln("auth publish   : ", authURL+"/publish")
		log.Debugln("auth subscribe : ", authURL+"/subscribe")
	} else {
		log.Warnln("auth : no auth url configured : bypassing!")
	}

	// Load CA cert
	authCACertPool := x509.NewCertPool()
	if authURL != "" && authCAFile != "" {
		caCert, err := ioutil.ReadFile(authCAFile)
		if err != nil {
			log.Fatal(err)
		}

		pemBlock, _ := pem.Decode(caCert)
		clientCert, err := x509.ParseCertificate(pemBlock.Bytes)
		if err != nil {
			log.Fatal(err)
		}

		clientCert.BasicConstraintsValid = true
		clientCert.IsCA = true
		clientCert.KeyUsage = x509.KeyUsageCertSign
		//clientCert.DNSNames = append(clientCert.DNSNames, "policy")
		authCACertPool.AddCert(clientCert)
		log.Println("auth using CA  : '" + authCAFile + "'")
	}

	tlsConfig := &tls.Config{RootCAs: authCACertPool}
	//tlsConfig.BuildNameToCertificate()
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
		//DisableCompression: true,
	}
	authClient = &http.Client{Transport: tr}

	if httpEnable || httpsEnable {
		wsMqttPrepare()
	}

	if httpEnable {
		go wsMqttListen()
	}

	if httpsEnable {
		go wssMqttListen()
	}

	if mqttEnable {
		go mqttListen()
	}

	if mqttsEnable {
		go mqttsListen()
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
