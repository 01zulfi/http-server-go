package handlers

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/request"
)

func handleFilesRoute(request request.Request, conn net.Conn, directory string) {
	slug := strings.TrimPrefix(request.Target, "/files/")
	filePath := directory + "/" + slug
	var response string
	switch request.Method {
	case "GET":
		contents, err := os.ReadFile(filePath)
		if err != nil {
			response = "HTTP/1.1 404 Not Found\r\n\r\n"
		} else {
			response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(contents), contents)
		}
	case "POST":
		body := request.Body
		bodyBytes := bytes.Trim([]byte(body), "\x00")
		err := os.WriteFile(filePath, bodyBytes, 0644)
		if err != nil {
			fmt.Println("Error while writing file: ", err.Error())
			response = "HTTP/1.1 500 Internal Server Error"
		} else {
			response = "HTTP/1.1 201 Created\r\n\r\n"
		}
	}
	handleResponse(response, conn)
}
