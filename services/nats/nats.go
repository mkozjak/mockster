package nats

import (
	"log"
	"net"
	"strconv"
)

func Start(port int, host string) error {
	log.Println("running nats mock server", port, host)

	ln, err := net.Listen("tcp", host+":"+strconv.Itoa(port))

	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			return err
		}

		go func(c net.Conn) {
		}(conn)
	}

	return nil
}
