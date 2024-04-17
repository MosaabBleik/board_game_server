package main

var (
	registerOption   = "register"
	unregisterOption = "unregister"
	moveOption       = "move"
)

type Player struct {
	ID       string `json:"id"`
	Position string `json:"position"`
	Rotation string `json:"rotation"`
}

type Message struct {
	Action string `json:"action"`
	Target string `json:"target"`
	Sender string `json:"sender"`
	Player Player `json:"player"`
}
