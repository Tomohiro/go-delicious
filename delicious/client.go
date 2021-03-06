package delicious

import (
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	apiEndpoint = "https://api.delicious.com/v1/"
)

// Client manages communication with the Delicious API.
type Client struct {
	// client provides request to API endpoints.
	client *http.Client

	// Endpint is the Delicious API endpoint.
	Endpoint string

	Posts *PostsService
}

// NewClient returns a new Delicious API client.
func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("access token is empty")
	}

	// Create an OAuth2 client to authentication.
	oauthClient := oauth2.NewClient(
		oauth2.NoContext,
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
	)

	c := &Client{
		client:   oauthClient,
		Endpoint: apiEndpoint,
	}
	c.Posts = &PostsService{client: c}

	return c, nil
}

// Do sends an HTTP request and returns HTTP response via http.Client.
//
// https://golang.org/pkg/net/http/#Client.Do
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
