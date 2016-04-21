package delicious

import (
	"encoding/xml"
	"net/http"
	"time"
)

// PostsService handles communication with post related
// methods of the Delicious API.
//
// Delicious API docs: https://github.com/domainersuitedev/delicious-api/blob/master/api/posts.md
type PostsService struct {
	client *Client
}

// Posts have multiple post.
type Posts struct {
	Posts *[]Post `xml:"post"`
}

// Post represents an bookmarked post.
//
// Delicious API docs: https://github.com/domainersuitedev/delicious-api/blob/master/api/posts.md#example
type Post struct {
	URL         string    `xml:"href,attr"`
	Description string    `xml:"description,attr"`
	Extended    string    `xml:"extended,attr"`
	Hash        string    `xml:"hash,attr"`
	Meta        string    `xml:"meta,attr"`
	Others      int       `xml:"others,attr"`
	Tag         string    `xml:"tag,attr"`
	Time        time.Time `xml:"time,attr"`
}

// Recent returns a list of the most recent posts
//
// Delicious API docs: https://github.com/domainersuitedev/delicious-api/blob/master/api/posts.md#v1postsrecent
func (s *PostsService) Recent() (*[]Post, error) {
	url := s.client.Endpoint + "/posts/recent"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	posts := &Posts{}
	if err = xml.NewDecoder(res.Body).Decode(&posts); err != nil {
		return nil, err
	}

	return posts.Posts, nil
}
