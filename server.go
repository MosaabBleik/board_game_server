package main

import (
	//"encoding/json"
	"errors"
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type Server struct {
	Name  string
	Conn  map[*websocket.Conn]bool
	Rooms map[string]*Room
}

func newServer() *Server {
	return &Server{
		Name:  "Server 1",
		Conn:  make(map[*websocket.Conn]bool),
		Rooms: make(map[string]*Room),
	}
}

func (s *Server) FindRoom(msg *Message) (*Room, error) {
	for k := range s.Rooms {
		if k == msg.Target {
			return s.Rooms[k], errors.New("Room exists already")
		}
	}
	return NewRoom(msg.Target, s), nil

}

func (s *Server) HandleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from a client")
	s.HandleRooms(ws)
}

func (s *Server) HandleRooms(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	var msg *Message
	var r *Room
	var client *Client
loop:
	for {
		n, err := ws.Read(buf)
		if err == io.EOF {
			break
		}
		msg = EncodeMessage(buf[:n])

		switch msg.Action {
		case registerOption:
			fmt.Println("Registered!!!")
			r, err = s.FindRoom(msg)
			if err != nil {
				fmt.Println(err)
			}

			client, err = r.FindClient(msg.Sender, ws)
			if err != nil {
				break loop
			}

			if err := r.RegisterClient(client); err == nil {
				fmt.Printf("Number of clients: %d\n", len(r.Clients))
				break loop
			}

		}
	}
	r.InformJoin(msg)
	s.HandleMessages(msg.Sender, ws, r, client)
}

func (s *Server) HandleMessages(clientName string, ws *websocket.Conn, r *Room, client *Client) {
	buf := make([]byte, 1024)

	defer func(clientName string, ws *websocket.Conn, r *Room) {
		ws.Close()
		isConnected := ws.IsClientConn()
		if !isConnected {
			fmt.Println("OPSSSSSSSSS")
		}
		delete(r.Clients, client)

		fmt.Println(r.Clients)
		//leaveMessage(clientName, r)
	}(clientName, ws, r)

loop:
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		msg := EncodeMessage(buf[:n])

		switch msg.Action {
		case registerOption:
			r.Broadcast([]byte(buf[:n]), msg.Sender)

		case unregisterOption:
			if err := r.UnregisterClient(msg, ws, s); err != nil {
				print("CLIENT DIDNT LEAVE")
				//ws.Write([]byte(fmt.Sprintf(`{"action": "unregister", "sender": "%s", "body": %s}`, msg.Sender, err.Error())))
			} else {
				fmt.Printf("Number of clients: %d\n", len(r.Clients))
				r.Broadcast([]byte(buf[:n]), msg.Sender)
				break loop
			}

		case moveOption:
			fmt.Println(string(buf[:n]))
			r.BroadcastMove([]byte(buf[:n]), msg.Sender)
		}
	}
}

// func leaveMessage(clientName string, r *Room) {
// 	errMsg := &Message{
// 		Action: "unregister",
// 		Sender: clientName,
// 	}

// 	b, err := json.Marshal(errMsg)
// 	if err != nil {
// 		fmt.Println("Something went wrong!")
// 	}
// 	r.Broadcast(b, )
// }
