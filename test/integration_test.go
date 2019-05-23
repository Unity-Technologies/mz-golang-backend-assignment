package test

import (
	"github.com/Applifier/golang-backend-assignment/internal/client"
	"github.com/Applifier/golang-backend-assignment/internal/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

const serverPort = 50000

func TestIntegration(t *testing.T) {
	srv := server.New()

	serverAddr := net.TCPAddr{Port: serverPort}
	require.NoError(t, srv.Start(&serverAddr))
	defer assertDoesNotError(t, srv.Stop)

	// Create clients
	client1 := createClientAndFetchID(t, 1)
	defer assertDoesNotError(t, client1.Close)
	client1Ch := make(chan client.IncomingMessage)
	defer close(client1Ch)

	client2 := createClientAndFetchID(t, 2)
	defer assertDoesNotError(t, client2.Close)
	client2Ch := make(chan client.IncomingMessage)
	defer close(client2Ch)

	client3 := createClientAndFetchID(t, 3)
	defer assertDoesNotError(t, client3.Close)
	client3Ch := make(chan client.IncomingMessage)
	defer close(client3Ch)

	t.Run("List other clients from each client", func(t *testing.T) {
		ids, err := client1.ListClientIDs()
		assert.NoError(t, err)
		assert.Equal(t, []uint64{2, 3}, ids)

		ids, err = client2.ListClientIDs()
		assert.NoError(t, err)
		assert.Equal(t, []uint64{1, 3}, ids)

		ids, err = client3.ListClientIDs()
		assert.NoError(t, err)
		assert.Equal(t, []uint64{1, 2}, ids)
	})

	t.Run("Send message from the first client to the two other clients", func(t *testing.T) {
		body := []byte("Hello world!")
		assert.Equal(t, nil, client1.SendMsg([]uint64{2, 3}, body))

		go client2.HandleIncomingMessages(client2Ch)
		incomingMessage := <-client2Ch
		assert.Equal(t, body, incomingMessage.Body)
		assert.Equal(t, uint64(1), incomingMessage.SenderID)

		go client3.HandleIncomingMessages(client3Ch)
		incomingMessage = <-client3Ch
		assert.Equal(t, body, incomingMessage.Body)
		assert.Equal(t, uint64(1), incomingMessage.SenderID)
	})
}

func assertDoesNotError(tb testing.TB, fn func() error) {
	assert.NoError(tb, fn())
}

func createClientAndFetchID(t *testing.T, expectedClientID uint64) *client.Client {
	cli := client.New()
	serverAddr := net.TCPAddr{Port: serverPort}
	require.NoError(t, cli.Connect(&serverAddr))
	id, err := cli.WhoAmI()
	assert.NoError(t, err)
	assert.Equal(t, expectedClientID, id)
	return cli
}
