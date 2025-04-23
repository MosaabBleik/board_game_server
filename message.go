package main

var (
	registerOption   = "register"
	unregisterOption = "unregister"
	moveOption       = "move"
	messageOption    = "message"
)

type Player struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	Number      string  `json:"number"`
	Position    string  `json:"position"`
	Rotation    string  `json:"rotation"`
	CurrentStep int     `json:"current_step"`
	StepsNumber int     `json:"steps_number"`
	Money       float64 `json:"money"`
}

type Building struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Owner string `json:"owner"`
}

type Message struct {
	Action string `json:"action"`
	Target string `json:"target"`
	Sender string `json:"sender"`
	Player Player `json:"player"`
}
