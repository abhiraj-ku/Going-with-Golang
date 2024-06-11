package main

import (
	"fmt"
	"net"
	"time"
)

func Check(dest string, port string) string {
	address := dest + ":" + port

	timeOut := time.Duration(5 * time.Second)

	conn, err := net.DialTimeout("tcp", address, timeOut)

	var status string

	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v", dest, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", dest,
			conn.LocalAddr(),
			conn.RemoteAddr())
	}
	return status
}
