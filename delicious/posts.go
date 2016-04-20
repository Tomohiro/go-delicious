package delicious

import (
	"encoding/xml"
	"net/http"
)

// PostsService handles communication with post related
// methods of the Delicious API.
//
// Delicious API docs: https://github.com/domainersuitedev/delicious-api/blob/master/api/posts.md
type PostsService struct {
	client *Client
}

// Post represents an bookmarked post.
//
// Delicious API docs: https://github.com/domainersuitedev/delicious-api/blob/master/api/posts.md#example
type Post struct {
	URL         string `xml:"href"`
	Description string `xml:"description"`
	Extended    string `xml:"extended"`
	Hash        string `xml:"hash"`
	Meta        string `xml:"meta"`
	Others      int    `xml:"others"`
	Tag         string `xml:"tag"`
	Time        string `xml:"time"`
}

// Recent returns recent posts.
//
// Delicious API docs: https://github.com/domainersuitedev/delicious-api/blob/master/api/posts.md#v1postsrecent
func (s *PostsService) Recent() (*[]Post, error) {
	url := s.client.Endpoint + "/posts/recent?"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	posts := new([]Post)

	if err = xml.NewDecoder(res.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return posts, nil
}
