package handlers

import (
	"fmt"
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/request"
)

func handleEchoRoute(request request.Request, conn net.Conn) {
	slug := strings.TrimPrefix(request.Target, "/echo/")
	acceptEncodingHeader := request.Headers["Accept-Encoding"]
	var response string

	if strings.Contains(acceptEncodingHeader, "gzip") {
		encoded, err := encodeGzip(slug)
		if err != nil {
			fmt.Println("Error while gzip encoding: ", err.Error())
			response = "HTTP/1.1 500 Internal Server Error"
		} else {
			response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Encoding: gzip\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(*encoded), *encoded)
		}
	} else {
		response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(slug), slug)
	}

	handleResponse(response, conn)
}
