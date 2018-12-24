package main


import (
	"fmt"
	"log"
	"net"
	"strconv"
)


func rawProtocolHandler(p *Protocol) {

	conn, err := net.Dial("tcp", p.Remote)
	if err != nil {
		log.Println(p.Ticket, err)
		sendResponseCallback(p, err.Error())
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%08d%s", len(p.Data), p.Data)

	lenBuffer := make([]byte, 8)

	if _, err := conn.Read(lenBuffer); err != nil {
		log.Println(p.Ticket, err)
		sendResponseCallback(p, err.Error())
		return
	}

	responseLen, err := strconv.Atoi(string(lenBuffer))
	if err != nil {
		log.Println(p.Ticket, err)
		sendResponseCallback(p, err.Error())
		return
	}

	responseBuffer := make([]byte, responseLen)

	if _, err := conn.Read(responseBuffer); err != nil {
		log.Println(p.Ticket, err)
		sendResponseCallback(p, err.Error())
		return
	}

	response := string(responseBuffer)

	sendResponseCallback(p, response)
}


