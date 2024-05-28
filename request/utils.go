package request

import (
	"net"
	"slices"
	"strings"
)

func GetRequestStringFromConnection(conn net.Conn) (*string, error) {
	buffer := make([]byte, 2048)
	_, err := conn.Read(buffer)
	if err != nil {
		return nil, err
	}
	requestString := string(buffer)
	return &requestString, nil
}

func ParseRequestString(request string) Request {
	splitBySpace := strings.Split(request, " ")
	splitByCrlf := strings.Split(request, "\r\n")
	method := splitBySpace[0]
	target := splitBySpace[1]
	rest := splitByCrlf[1:]
	body := rest[len(rest)-1]
	headers := slices.Delete(rest, len(rest)-1, len(rest))

	headerMap := make(map[string]string)
	for _, header := range headers {
		if strings.Contains(header, ": ") {
			split := strings.Split(header, ": ")
			headerMap[split[0]] = split[1]
		}
	}

	return Request{Method: method, Target: target, Headers: headerMap, Body: body}
}
