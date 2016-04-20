package delicious

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client, _ = NewClient("5095047-5ddd35cd5ad5ad086436b8c6bafb84c2")
	//client.Endpoint = server.URL
}

func teardown() {
	server.Close()
}

func TestRecent(t *testing.T) {
	setup()
	defer teardown()

	_, err := client.Posts.Recent()
	if err != nil {
		t.Fatalf("Recent returned error: %v", err)
	}
}
