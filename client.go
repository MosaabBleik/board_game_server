package main

import "golang.org/x/net/websocket"

type Client struct {
	Name string
	Conn *websocket.Conn
}

func NewClient(name string, ws *websocket.Conn) *Client {
	return &Client{
		Name: name,
		Conn: ws,
	}
}
