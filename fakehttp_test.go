package fakehttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/couchbaselabs/go.assert"
)

func TestUsage(t *testing.T) {

	// startup fake server
	testServer := NewHTTPServerWithPort(4444)
	testServer.Start()

	// tell it to respond with 200 status and this fake response
	// on the next request it receives
	fakeResponse := "<html>foo</html>"
	testServer.Response(200, nil, fakeResponse)

	// send a request to the fake server and read response
	urlString := fmt.Sprintf("%s/foo.html", testServer.URL.String())
	res, err := http.Get(urlString)
	if err != nil {
		panic("unexpected error")
	}

	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("unexpected error")
	}

	// make sure that the response is our fake response
	assert.True(t, string(responseBody) == fakeResponse)

}

func TestSavedRequestConcurrency(t *testing.T) {
	testServer := NewHTTPServerWithPort(4445)
	testServer.Start()
	testServer.Response(200, nil, "OK")

	go func() {
		for i := 0; i < 100; i++ {
			urlString := fmt.Sprintf("%s/foo.html", testServer.URL.String())
			if _, err := http.Get(urlString); err != nil {
				panic("unexpected error")
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			// safe access to savedRequest slice:
			_ = testServer.Requests()
			// Causes race:
			// _ = testServer.SavedRequests
		}
	}()

	time.Sleep(time.Millisecond)
}
