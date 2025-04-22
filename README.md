# Rich Errors (richerr)

This library provides a very small set of enhancements on
top of the default `error` type in Go to allow arbitrary
metadata to be included on errors in a way that is similar
to the way structured logging packages behave. It is common
for errors to eventually end up being passed to a structured
logger of some sort, so having error metadata that is easy
to attach to a log message makes sense.

There are no dependencies other than the standard library
and an assertion library for the unit tests.

For example, we can attach parameter values to an error to
help us track down the cause later.

```go
func FetchUser(id string) (User, error) {
    // ... whatever setup we need 
    
    resp, err := db.Query(...)
    if err != nil {
        return nil, richerr.Wrap(err, "failed to fetch user").
            WithField("user_id", id)
    }
	
    // ... happy path
}
```

Errors can also be associated with "scopes" to allow field names
to be recovered later in a way that allows fields with the same
name on different errors to be distinguished. The example below
isn't necessarily a good idea, the fields could simply be named
more clearly, but it's the type of thing that might happen in a
large team where different people are implementing different parts
of the system.

```go
func GetPosts(userId string) ([]Post, error) {
	// ... get the user's posts
	
	for _, postId := range userPosts {
		post, err := GetPost(postId)
        if err != nil {
		return nil, richerr.Wrap(err, "failed to fetch post").
			WithScope("GetPosts").
			WithField("id", userId)
        }
    }
	
    return posts, nil
}

func GetPost(postId string) (*Post, error) {
    // ... try to find the post

    return nil, richerr.New("no post found").
		WithScope("GetPost").
		WithField("id", postId)
}
```

Once the error has bubbled up to the spot where you want to handle it,
any fields introduced along the way can be recovered with the `Collect`
function.

```go
fields := richerr.Collect(err)

// fields now contains a slice of all fields added to the tree
// of errors rooted at err
```
