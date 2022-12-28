package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bmcculley/mockexample/mocks"
)

func init() {
	Client = &mocks.MockClient{}
}

func TestGet(t *testing.T) {
	html := `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Hello World!</title>
</head>
<body>
	<h1>Just a test page</h1>
</body>
</html>`
	r := ioutil.NopCloser(bytes.NewReader([]byte(html)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	request, err := Get("http://static.ex.net/extension-testing/")
	if err != nil {
		t.Errorf("Got %q want nil", err)
	}

	var reader io.ReadCloser
	reader = request.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	got := buf.String()

	if got != html {
		t.Errorf("got %q want %q", got, html)
	}
}