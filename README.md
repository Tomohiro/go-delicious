go-delicious
================================================================================

go-delicious is a library for Go for accessing the [Delicous API][]

[Delicious API]: https://github.com/domainersuitedev/delicious-api


Installation
--------------------------------------------------------------------------------

Use `go get`:

```sh
$ go get -d github.com/Tomohiro/go-delicious
```


Usage
--------------------------------------------------------------------------------

### Create a client to accessing the Delicious API

Import this package like this:

```go
import "github.com/Tomohiro/go-delicious/delicious"
```

Create a client with your [Delicious application credential](https://delicious.com/oauth/applications):

```go
delicious, err := delicious.NewClient("your access token")
if err != nil {
	panic(err)
}
```

### Posts

#### Fetch all bookmarks

```go
list, _ := delicious.Posts.All(&delicious.FetchOptions{Start: 1, Resutls: 50})
for _, post := range *list.Posts {
	fmt.Println(post.URL) // http://www.example.com/
}
```

#### Get one or more bookmarks

```go
list, _ := delicious.Posts.Get(&delicious.FetchOptions{Tag: "Programming"})
for _, post := range *list.Posts {
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
