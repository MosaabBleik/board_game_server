package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/net/websocket"
)

type Room struct {
	Name    string
	Clients map[*Client]bool
}

func NewRoom(roomName string, s *Server) *Room {
	r := &Room{
		Name:    roomName,
		Clients: make(map[*Client]bool),
	}
	s.Rooms[roomName] = r
	return r
}

func RemoveRoom(roomName string, s *Server) {
	delete(s.Rooms, roomName)
	fmt.Println("Number of rooms:::", len(s.Rooms))
}

func (r *Room) InformJoin(msg *Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Couldn't parse the message")
	}
	r.Broadcast([]byte(data), msg.Sender)
}

func (r *Room) Broadcast(b []byte, sender string) {
	for c := range r.Clients {
		go SendMessage(c, b)
	}
}

func (r *Room) BroadcastMove(b []byte, sender string) {
	for c := range r.Clients {
		if c.Name != sender {
			go SendMessage(c, b)
		}
	}
}

func SendMessage(c *Client, b []byte) {
	if _, err := c.Conn.Write(b); err != nil {
		fmt.Println("write error:", err)
	}
}

func welcomeMessage(ws *websocket.Conn) {
	ws.Write([]byte("welcome in the club"))
}

func EncodeMessage(data []byte) *Message {
	msg := &Message{}
	if err := json.Unmarshal(data, msg); err != nil {
		return nil
	}
	return msg
}

func (r *Room) RegisterClient(c *Client) error {
	_, ok := r.Clients[c]
	if ok {
		return errors.New("Client exists already")
	}

	r.Clients[c] = true
	return nil
}

func (r *Room) UnregisterClient(msg *Message, ws *websocket.Conn, s *Server) error {
	clientName := msg.Sender
	for c := range r.Clients {
		if c.Name == clientName {
			delete(r.Clients, c)
			if len(r.Clients) == 0 {
				RemoveRoom(r.Name, s)
			}
			return nil
		}
	}
	return errors.New("client doesn't exist")
}

func (r *Room) FindClient(clientName string, ws *websocket.Conn) (*Client, error) {
	for k := range r.Clients {
		if k.Name == clientName {
			return k, errors.New("client exists already")
		}
	}
	return NewClient(clientName, ws), nil
}

// func (r *Room) DeleteClient(clientName string, s *Server) error {
// 	for c := range r.Clients {
// 		if c.Name == clientName {
// 			delete(r.Clients, c)
// 			if len(r.Clients) == 0 {
// 				RemoveRoom(r.Name, s)
// 			}
// 			return nil
// 		}
// 	}
// 	return errors.New("client doesn't exist")
// }
