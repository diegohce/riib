package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func newTicket(p *Protocol) string {

	var rv string

	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		log.Println("newTicket::", err)
		rv = newTicket2(p)

	} else {

		rv = base64.StdEncoding.EncodeToString(b)
	}

	return rv
}

func newTicket2(p *Protocol) string {

	h := sha256.New()

	hashData := fmt.Sprintf("%+v", p)

	h.Write([]byte(hashData))

	rv := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return rv
}
