# Rich Errors (richerr)

This library provides a very small set of enhancements on
top of the default `error` type in Go to allow arbitrary
metadata to be included on errors in a way that is similar
to the way structured logging packages behave. It is common
for errors to eventually end up being passed to structured
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
