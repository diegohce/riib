package main

// Protocol : spec between requests and calls
type Protocol struct {
	Remote      string `json:"remote"`
	Data        string `json:"data"`
	DataType    string `json:"data_type"`
	CallbackURL string `json:"callback_url"`
	Method      string `json:"method"`
	Ticket      string
}

type protocolHandler func(*Protocol)

var (
	protocols map[string]protocolHandler
)

func init() {
	protocols = map[string]protocolHandler{
		"http": httpProtocolHandler,
		"raw":  rawProtocolHandler,
	}
}
