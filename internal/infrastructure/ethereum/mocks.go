package ethereum

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

type MockBody struct {
	io.Reader
}

func (m *MockBody) Close() error {
	return nil
}

func NewMockBody(content string) *MockBody {
	return &MockBody{Reader: ioutil.NopCloser(strings.NewReader(content))}
}
