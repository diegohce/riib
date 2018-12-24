package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func httpProtocolHandler(p *Protocol) {

	client := http.Client{}

	method := p.Method
	if method == "" {
		method = "POST"
	}

	var response *http.Response

	switch method {

		case "GET":
		{
			var err error

			response, err = client.Get(p.Remote)
			if err != nil {
				log.Println(p.Ticket, err)
				sendResponseCallback(p, err.Error())
				return
			}
		}
		default:
		{
			var err error

			buf := bytes.NewReader([]byte(p.Data))
			response, err = client.Post(p.Remote, "application/json", buf)
			if err != nil {
				log.Println(p.Ticket, err)
				sendResponseCallback(p, err.Error())
				return
			}
		}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(p.Ticket, err)
		sendResponseCallback(p, err.Error())
		return
	}
	defer response.Body.Close()

	sendResponseCallback(p, string(body))
}
