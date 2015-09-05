package Splunk

import (
	"net"

	"../../../src"
	. "../../utility"
)

type Splunk struct {
	Connection *net.TCPConn // Connection
}

func SplunkInit() *Splunk {

	ServerAddr, err := net.ResolveTCPAddr(config.SPLUNK_PROTOCOL, config.SPLUNK_ADDR)
	ErrCheck(err)

	Conn, err := net.DialTCP(config.SPLUNK_PROTOCOL, nil, ServerAddr)
	ErrCheck(err)

	return &Splunk{Connection: Conn}
}

func (Conn *Splunk) Close() {
	Conn.Close()
}
