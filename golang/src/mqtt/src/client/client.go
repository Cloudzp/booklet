package client

import (
	"bytes"
	"fmt"
	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
	"net"
)

type  Client struct {
	addr string
	exit chan struct{}
	conn *mqtt.ClientConn

	TQS []proto.TopicQos
}

func NewClient(addr string, tqs ...proto.TopicQos) (*Client, error){

	conn , err := net.Dial("tcp",addr)
	if err != nil {
		return nil, err
	}
	c := &Client{
		exit: make(chan  struct{} ),
		conn: mqtt.NewClientConn(conn),
        TQS:tqs,
	}

	if err := c.conn.Connect("",""); err != nil {
		return nil, err
	}

	return c,nil
}

func (c *Client)Run(){

	c.conn.Subscribe(c.TQS)
	go func() {
		var buf =  &bytes.Buffer{}

		for msg := range c.conn.Incoming{
			err := msg.Payload.WritePayload(buf)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("topic:", msg.TopicName)
			fmt.Println("payload:", buf.String())
			buf.Reset()
		}
	}()
	<- c.exit
}

func (c *Client)Stop(){
	 c.exit <- struct{}{}
	 c.conn.Disconnect()
}
