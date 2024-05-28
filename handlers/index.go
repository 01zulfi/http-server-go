package handlers

import "net"

func handleIndexRoute(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\n\r\n"
	handleResponse(response, conn)
}
