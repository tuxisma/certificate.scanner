package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	conn, err := tls.Dial("tcp", "tuxisma.github.io:443", nil)
	if err != nil {
		panic("Server does not support SSL certificate err: " + err.Error())
	}

	err = conn.VerifyHostname("tuxisma.github.io")
	if err != nil {
		panic("Hostname does not match with certificate: " + err.Error())
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))
}
