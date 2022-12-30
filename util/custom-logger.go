package util

import (
	"bytes"
	"net/http"
	"os"
)

type NetworkLogger struct{}

func (n NetworkLogger) Write(b []byte) (int, error) {
	http.Post(os.Getenv("NETLOG_URL"), "", bytes.NewReader(b))
	return 1, nil
}