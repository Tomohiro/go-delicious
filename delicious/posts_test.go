package delicious

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client, _ = NewClient("DUMMY_ACCESS_TOKEN")
	client.Endpoint = server.URL
}

func teardown() {
	server.Close()
}

func TestRecent(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/posts/recent", func(w http.ResponseWriter, r *http.Request) {
		// Set response headers
		w.Header().Set("Content-Type", "application/xml")

		// Set 200 OK as HTTP status code.
		w.WriteHeader(http.StatusOK)

		// Set response body
		fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?>
			<posts tag="" user="johndoe">
				<post
					description="Post Description"
					extended=""
					hash="bb86b8f48fef41b2593ad53b797d7de6"
					href="http://example.com/post"
					private="no"
					shared="yes"
					tag="test dummy"
					time="2016-04-21T05:21:27Z"
				/>
			</posts>
		`)
	})

	posts, err := client.Posts.Recent()

	if err != nil {
		t.Fatalf("Recent returned error: %v", err)
	}

	actual := posts
	expectedTime, _ := time.Parse(time.RFC3339, "2016-04-21T05:21:27Z")
	expected := &[]Post{{
		URL:         "http://example.com/post",
		Description: "Post Description",
		Tag:         "test dummy",
		Hash:        "bb86b8f48fef41b2593ad53b797d7de6",
		Extended:    "",
		Others:      0,
		Meta:        "",
		Time:        expectedTime,
	}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("List returned %+v, want %+v", actual, expected)
	}
}
