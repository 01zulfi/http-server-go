package handlers

import (
	"fmt"
	"net"
)

func handleResponse(response string, conn net.Conn) {
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response: ", err.Error())
	}
	err = conn.Close()
	if err != nil {
		fmt.Println("Error closing connection: ", err.Error())
	}
}
