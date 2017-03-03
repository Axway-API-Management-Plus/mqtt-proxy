package main

import (
	"net"
	"strconv"
	"sync"
	"sync/atomic"

	log "github.com/Sirupsen/logrus"
	uuid "github.com/satori/go.uuid"
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

func (session *Session) forwardHalf(c1 net.Conn, c2 net.Conn) {
	defer c1.Close()
	defer c2.Close()
	defer session.wg.Done()
	//io.Copy(c1, c2)

	for {
		log.Println("Session", session.id, "- Wait Packet", c1.RemoteAddr().String(), c2.RemoteAddr().String())
		err := session.ForwardMQTTPacket(c1, c2)
		if err != nil {
			session.closed = true
			break
		}
	}
}

func (session *Session) Stream(conn net.Conn) {
	session.inbound = conn
	session.wg.Add(2)
	addr := mqttHost + ":" + strconv.Itoa(mqttPort)
	log.Println("Session", session.id, "- Dialing...", session.id, addr)
	client, err := net.Dial("tcp", addr)
	if err != nil {
		log.Errorln("Session", session.id, "- Dial failed :", addr, err)
		return
	}
	log.Println("Session", session.id, "- Connected", conn.RemoteAddr().String(), addr)
	session.outbound = client
	go session.forwardHalf(client, conn)
	go session.forwardHalf(conn, client)
	session.wg.Wait()

	atomic.AddInt32(&globalSessionCount, -1)
	log.Println("Session", session.id, "Closed", conn.LocalAddr().String(), globalSessionCount)
}
