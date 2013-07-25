# fakehttp

Fake in-process webserver for unit testing code which depends on an external webserver.

Fakehttp is trying to accomplish the same goal as the Ruby fakeweb library, but at the network layer rather than via monkey patching.  

![architecture.png](http://cl.ly/image/1n0n3y283z0a/Screen%20Shot%202013-07-24%20at%2010.02.57%20PM.png)

Essentially your tests just need to send requests to the fakehttp server, which allows you to control exactly what responses it returns.

# fakehttp vs goamz.testutil

This is a fork of the `testutil` component of Gustavo Niemeyer's [goamz](https://github.com/soundcloud/goamz) library with the following changes:

* Extracted into dedicated repo
* Package rename testutil -> fakehttp
* README documentation
* Unit Test

# Usage

See the `TestUsage` method in the unit test for a usage example

```go
var testServer = NewHTTPServer()
testServer.Start()

testServer.Response(200, nil, "<html>foo</html>")

res, err := http.Get("%s/foo.html", testServer.URL)
responseBody, err := ioutil.ReadAll(res.Body)

assert.True(t, string(responseBody) == "<html>foo</html>")
```


# Running unit tests

```
$ git clone ${this_repo}
$ cd fakehttp
$ go build -v && go test
```

# References

* [Golang group thread](https://groups.google.com/forum/#!topic/golang-nuts/6AN1E2CJOxI) regarding unit testing with http server dependencies

* [goamz](https://github.com/soundcloud/goamz) - where this code originally came from