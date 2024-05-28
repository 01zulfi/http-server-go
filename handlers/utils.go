package handlers

import (
	"bytes"
	"compress/gzip"
)

func encodeGzip(s string) (*string, error) {
	var b bytes.Buffer
	enc := gzip.NewWriter(&b)
	_, err := enc.Write([]byte(s))
	if err != nil {
		return nil, err
	}
	err = enc.Close()
	if err != nil {
		return nil, err
	}
	stringified := b.String()
	return &stringified, nil
}
