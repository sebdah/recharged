package handlers_test

import (
	"io"
	"net/http/httptest"
)

var (
	server *httptest.Server
	reader io.Reader
)
