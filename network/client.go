package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Address string
	packer  *NormalPacker
}

func NewClient(address string) *Client {
	client := &Client{
		Address: address,
		packer:  NewNormalPacker(binary.BigEndian),
	}
	return client
}

func (c *Client) Run() {
	// tcp6 -> IPV6-only ?
	conn, err := net.Dial("tcp6", c.Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	go c.Write(conn)
	go c.Read(conn)

}

func (c *Client) send(conn net.Conn, message *Message) {
	err := conn.SetWriteDeadline(time.Now().Add(time.Second))

	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write(bytes)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *Client) Write(conn net.Conn) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			c.send(conn, &Message{
				Id:   111,
				Data: []byte("client say: Hello World!"),
			})
		}
	}
}

func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.UnPack(conn)
		if _, ok := err.(net.Error); ok {
			fmt.Println(err)
			continue
		}
		fmt.Println("resp message: ", string(message.Data))
	}
}
