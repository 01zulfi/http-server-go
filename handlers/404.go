package handlers

import "net"

func handle404Route(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\n\r\n"
	handleResponse(response, conn)
}
