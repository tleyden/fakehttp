package fakehttp

import (
	"fmt"
	"github.com/couchbaselabs/go.assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUsage(t *testing.T) {

	// startup fake server
	testServer := NewHTTPServer()
	testServer.Start()

	// tell it to respond with 200 status and this fake response
	// on the next request it receives
	fakeResponse := "<html>foo</html>"
	testServer.Response(200, nil, fakeResponse)

	// send a request to the fake server and read response
	urlString := fmt.Sprintf("%s/foo.html", testServer.URL)
	res, err := http.Get(urlString)
	if err != nil {
		panic("unexpected error")
	}
	defer responseBody.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("unexpected error")
	}

	// make sure that the response is our fake response
	assert.True(t, string(responseBody) == fakeResponse)

}
