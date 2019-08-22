package client

import (
	"fmt"
	"net"
)

type IncomingMessage struct {
	SenderID uint64
	Body     []byte
}

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (cli *Client) Connect(serverAddr *net.TCPAddr) error {
	fmt.Println("TODO: Connect to the server using the given address")
	return nil
}

func (cli *Client) Close() error {
	fmt.Println("TODO: Close the connection to the server")
	return nil
}

func (cli *Client) WhoAmI() (uint64, error) {
	fmt.Println("TODO: Fetch the ID from the server")
	return 0, nil
}

func (cli *Client) ListClientIDs() ([]uint64, error) {
	fmt.Println("TODO: Fetch the IDs from the server")
	return nil, nil
}

func (cli *Client) SendMsg(recipients []uint64, body []byte) error {
	fmt.Println("TODO: Send the message to the server")
	return nil
}

func (cli *Client) HandleIncomingMessages(writeCh chan<- IncomingMessage) {
	fmt.Println("TODO: Handle the messages from the server")
}
