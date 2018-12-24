package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	protocol := ps.ByName("protocol")

	handler, ok := protocols[protocol]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid protocol '%s'", protocol)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error procesing request %s", err)
		return
	}
	defer r.Body.Close()

	var p Protocol

	if err := json.Unmarshal(body, &p); err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error procesing request %s", err)
		return
	}

	//ticket := "ASDQWEERTFGHHJTYUGHJ123345567"
	ticket := newTicket(&p)

	p.Ticket = ticket

	go handler(&p)

	fmt.Fprintf(w, "%s", ticket)
}

func main() {

	router := httprouter.New()

	router.POST("/riib/:protocol/run", requestHandler)

	log.Fatal(http.ListenAndServe(":9999", router))

}
