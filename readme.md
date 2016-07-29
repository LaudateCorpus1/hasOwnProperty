# hasOwnProperty [![Build Status](https://travis-ci.org/WatchBeam/hasOwnProperty.svg?branch=master)](https://travis-ci.org/WatchBeam/hasOwnProperty) [![godoc reference](https://godoc.org/github.com/WatchBeam/hasOwnProperty?status.png)](https://godoc.org/github.com/WatchBeam/hasOwnProperty)

Go's JSON package does not provide a way to distinguish between `null` keys and keys which do not exist on a structure. This package provides such functionality.

```go
const json = []byte(`{"hello": {"world": true}}`)


ok := hasOwnProperty.Test(json, "hello") // true
ok = hasOwnProperty.Test(json, "hello.world") // true
ok = hasOwnProperty.Test(json, "goodbye") // false
```
