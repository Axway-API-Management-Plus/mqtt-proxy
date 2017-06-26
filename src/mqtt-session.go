package main

import (
	"crypto/tls"
	"net"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Session struct {
	id string
	//id int64
	wg sync.WaitGroup
	//deviceName string
	Username         string
	ClientIdentifier string
	inbound          net.Conn
	outbound         net.Conn
	closed           bool
}

var globalSessionID int64
var globalSessionCount int32

func NewSession() *Session {
	var session Session
	//g := atomic.AddInt64(&globalSessionID, 1)
	session.id = uuid.NewV4().String()
	atomic.AddInt32(&globalSessionCount, 1)
	return &session
}

func (session *Session) forwardHalf(way string, c1 net.Conn, c2 net.Conn) {
	defer c1.Close()
	defer c2.Close()
	defer session.wg.Done()
	//io.Copy(c1, c2)

	for {
		log.Println("Session", session.id, way, "- Wait Packet", c1.RemoteAddr().String(), c2.RemoteAddr().String())
		err := session.ForwardMQTTPacket(way, c1, c2)
		if err != nil {
			session.closed = true
			break
		}
	}
}

func (session *Session) DialOutbound() error {
	addr := mqttBrokerHost + ":" + strconv.Itoa(mqttBrokerPort)
	log.Println("Session", session.id, "- Dialing...", session.id, addr)
	client, err := net.Dial("tcp", addr)
	if err != nil {
		log.Errorln("Session", session.id, "- Dial failed :", addr, err)
		return err
	}
	log.Println("Session", session.id, "- Connected", session.inbound.RemoteAddr().String(), addr)
	session.outbound = client
	return nil
}

func (session *Session) Stream(conn net.Conn) {
	session.inbound = conn
	session.wg.Add(2)
	err := session.DialOutbound()
	if err != nil {
		return
	}
	go session.forwardHalf("<", session.outbound, session.inbound)
	go session.forwardHalf(">", session.inbound, session.outbound)
	session.wg.Wait()

	atomic.AddInt32(&globalSessionCount, -1)
	log.Println("Session", session.id, "Closed", conn.LocalAddr().String(), globalSessionCount)
}

func mqttAccept(l net.Listener) {
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

func mqttListen() {
	// Listen for incoming connections.
	addr := mqttHost + ":" + strconv.Itoa(mqttPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("mqtt: Error listening mqtt://"+addr, err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Println("mqtt: listening on mqtt://" + addr)

	mqttAccept(l)
}

func mqttsListen() {
	// Listen for incoming connections.
	cert, err := tls.LoadX509KeyPair(mqttsCert, mqttsKey)
	if err != nil {
		log.Fatalf("mqtts: loadkeys: %s", err)
		os.Exit(1)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	addr := mqttsHost + ":" + strconv.Itoa(mqttsPort)
	l, err := tls.Listen("tcp", addr, &config)
	if err != nil {
		log.Println("mqtts: Error listening mqtts://"+addr, err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Println("mqtts: listening on mqtts://" + addr)

	mqttAccept(l)
}
