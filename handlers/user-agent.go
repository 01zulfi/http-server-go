package handlers

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/request"
)

func handleUserAgentRoute(request request.Request, conn net.Conn) {
	userAgentHeader := request.Headers["User-Agent"]
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(userAgentHeader), userAgentHeader)
	handleResponse(response, conn)
}
