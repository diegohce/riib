package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func sendResponseCallback(p *Protocol, response string) {

	myResponse := map[string]string{p.Ticket: response}

	j, err := json.Marshal(myResponse)
	if err != nil {
		log.Println(p.Ticket, err)
		return
	}

	buf := bytes.NewReader(j)

	cbResponse, err := http.Post(p.CallbackUrl, "application/json", buf)
	if err != nil {
		log.Println(p.Ticket, err)
		return
	}
	defer cbResponse.Body.Close()
}
