package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)


func newTicket(p *Protocol) string {

	h := sha256.New()

	hashData := fmt.Sprintf("%+v", p)

	h.Write([]byte(hashData))

	rv := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return rv
}

