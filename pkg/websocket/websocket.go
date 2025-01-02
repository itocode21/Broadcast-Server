package websocket

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

// Client структура для представления клиента
type Client struct {
	Conn *websocket.Conn
}

// NewClient создает нового клиента
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn: conn,
	}
}

func Dial(address string) (*Client, error) {
	u := url.URL{Scheme: "ws", Host: address, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), http.Header{})
	if err != nil {
		return nil, err
	}
	client := NewClient(conn)
	return client, nil
}
