package gpntron

import (
	"log"
	"net"
)

var _ Executor = &executor{}

type Executor interface {
	Join(username, password string)
	Move(move Move)
	Chat(message string)
}

type executor struct {
	conn     *net.TCPConn
	Receiver Receiver
}

func (c *executor) Join(username, password string) {
	_, err := c.conn.Write(serialize(Join, username, password))
	if err != nil {
		log.Println(err)
	}
}

func (c *executor) Move(move Move) {
	_, err := c.conn.Write(serialize(Moving, move))
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Sent move", move)
	}
}

func (c *executor) Chat(message string) {
	_, err := c.conn.Write(serialize(Message, message))
	if err != nil {
		log.Println(err)
	}
}
