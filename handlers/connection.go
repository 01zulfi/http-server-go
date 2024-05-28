package handlers

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/request"
)

func HandleConnection(conn net.Conn, directoryFlag string) {
	requestString, err := request.GetRequestStringFromConnection(conn)

	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
		os.Exit(1)
	}

	request := request.ParseRequestString(*requestString)
	fmt.Println("Request received: ", request)

	switch {
	case request.Target == "/":
		handleIndexRoute(conn)
	case strings.HasPrefix(request.Target, "/echo"):
		handleEchoRoute(request, conn)
	case strings.HasPrefix(request.Target, "/user-agent"):
		handleUserAgentRoute(request, conn)
	case strings.HasPrefix(request.Target, "/files"):
		handleFilesRoute(request, conn, directoryFlag)
	default:
		handle404Route(conn)
	}
}
