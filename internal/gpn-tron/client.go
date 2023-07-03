package gpntron

import (
	"bufio"
	"log"
	"net"
	"sync"
	"time"
)

type ClientOptions struct {
	Host               string
	Username, Password string
	TimeoutDuration    time.Duration
}

func NewClient(options ClientOptions) Client {
	return &client{
		done:          false,
		ClientOptions: options,
	}
}

type Client interface {
	Run(Receiver)
	Disconnect()
}

var _ Client = &client{}

type client struct {
	done bool
	ClientOptions
}

func (c *client) Disconnect() {
	c.done = true
	log.Println("Disconnecting")
}

func (c *client) Run(receiver Receiver) {

	var wg sync.WaitGroup

	tcpAddr, err := net.ResolveTCPAddr("tcp", c.Host)
	if err != nil {
		log.Fatal("ResolveTCPAddr failed:", err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Dial failed:", err.Error())
	}

	defer conn.Close()

	ex := &executor{
		conn:     conn,
		Receiver: receiver,
	}

	wg.Add(1)

	go func() {

		defer wg.Done()

		bufReader := bufio.NewReader(conn)

		for !c.done {
			// Set a deadline for reading. Read operation will fail if no data
			// is received after deadline.
			conn.SetReadDeadline(time.Now().Add(c.TimeoutDuration))

			// Read tokens until delimeter occurs
			bytes, err := bufReader.ReadBytes(PackageEnd)
			if err != nil {
				log.Println(err)
				break
			}

			command := string(bytes)[:len(bytes)-1]

			execute(command, receiver, ex)
		}
	}()

	// begin with the protocoll
	ex.Join(c.Username, c.Password)

	wg.Wait()
	log.Print("Client finished")
}
