go-delicious
================================================================================

[![Build Status](https://img.shields.io/travis/Tomohiro/go-delicious.svg?style=flat-square)](https://travis-ci.org/Tomohiro/go-delicious)
[![Coverage Status](https://img.shields.io/coveralls/Tomohiro/go-delicious.svg?style=flat-square)](https://coveralls.io/github/Tomohiro/go-delicious)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/Tomohiro/go-delicious/blob/master/LICENSE)

go-delicious is a library for Go for accessing the [Delicious API][]

[Delicious API]: https://github.com/domainersuitedev/delicious-api


Installation
--------------------------------------------------------------------------------

Use `go get`:

```sh
$ go get -d github.com/Tomohiro/go-delicious
```


Usage
--------------------------------------------------------------------------------

### Get an access token

Create a your application at [Delicious application credential](https://delicious.com/oauth/applications) and get an access token:

```sh
$ curl -X POST "https://avosapi.delicious.com/api/v1/oauth/token?client_id=$DELICIOUS_CLIENT_ID&client_secret=$DELICIOUS_CLIENT_SECRET&grant_type=credentials&username=$DELICIOUS_USERNAME&password=$DELICIOUS_PASSWORD" | jq -r '.access_token'
1936952-4dad55cd5ad5ad083942b8c6bafb24a9
```

### Create a client to accessing the Delicious API

Import this package like this:

```go
import "github.com/Tomohiro/go-delicious/delicious"
```

Create a client with your access token:

```go
delicious, err := delicious.NewClient("your access token")
if err != nil {
	panic(err)
}
```

### Posts

#### Fetch all bookmarks

```go
posts, _ := delicious.Posts.All(&delicious.FetchOptions{Start: 1, Results: 50})
for _, post := range *posts {
	fmt.Println(post.URL) // http://www.example.com/
}
```

#### Get one or more bookmarks

```go
posts, _ := delicious.Posts.Get(&delicious.FetchOptions{Tag: "Programming"})
for _, post := range *posts {
	fmt.Println(post.URL) // http://www.example.com/
```

#### Add a new bookmark

```go
post, _ := delicious.Posts.Add(&delicious.Bookmark{URL: "http://example.com"})
fmt.Println(post.URL) // http://www.example.com/
```

#### Delete

```go
result, _ := delicious.Posts.Delete("http://www.example.com")
```


### Tags

#### Fetch all tags

```go
tags, _ := delicious.Tags.All()
for _, tag := range *tags {
	fmt.Println(tag.Name) // Programming
}
```


#### Delete a tag from all posts

```go
result := delicious.Tags.Delete("Programming")
```


#### Rename a tag on all posts

```go
result, _ := delicious.Tags.Rename("Programming", "Coding")
```


Contributing
--------------------------------------------------------------------------------

Please check out the [CONTIRBUTING](CONTRIBUTING.md) guideline.


LICENSE
--------------------------------------------------------------------------------

&copy; 2016 Tomohiro TAIRA.

This project is licensed under the MIT license. See [LICENSE](LICENSE) for details.
