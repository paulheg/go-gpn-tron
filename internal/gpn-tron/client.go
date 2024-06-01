package gpntron

import (
	"bufio"
	"context"
	"fmt"
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
		ClientOptions: options,
	}
}

type Client interface {
	Run(context.Context, Receiver) error
}

var _ Client = &client{}

type client struct {
	ClientOptions
}

func (c *client) Run(ctx context.Context, receiver Receiver) error {

	var wg sync.WaitGroup

	tcpAddr, err := net.ResolveTCPAddr("tcp", c.Host)
	if err != nil {
		return fmt.Errorf("ResolveTCPAddr failed: %w", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return fmt.Errorf("Dial failed %w:", err)
	}

	// Set a deadline for reading. Read operation will fail if no data
	// is received after deadline.
	conn.SetReadDeadline(time.Now().Add(c.TimeoutDuration))

	defer conn.Close()

	ex := &executor{
		conn:     conn,
		Receiver: receiver,
	}

	commandChannel := make(chan string)

	// consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				conn.Close()
				return
			case command := <-commandChannel:
				execute(command, receiver, ex)
			}
		}
	}()

	// generator
	wg.Add(1)
	go func() {
		defer wg.Done()

		bufReader := bufio.NewReader(conn)

		for {
			// Read tokens until delimeter occurs
			bytes, err := bufReader.ReadBytes(PackageEnd)
			if err != nil {
				log.Printf("error reading new commands: %s", err)
				break
			}

			command := string(bytes)[:len(bytes)-1]
			commandChannel <- command
		}
	}()

	// begin with the protocoll
	ex.Join(c.Username, c.Password)

	wg.Wait()
	log.Print("Client finished")
	return nil
}
