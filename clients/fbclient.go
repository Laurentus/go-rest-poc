package clients

import "net/http"
import "log"

type FacebookResponseBody map[string]string

type httpInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type FacebookClient struct {
	uri    string
	client httpInterface
}

func NewFacebookClient(http httpInterface) FacebookClient {
	return FacebookClient{"https://graph.facebook.com/", http}
}

func (fc FacebookClient) Uri(path string) string {
	return fc.uri + path
}

func (fc FacebookClient) fetchRequest(c chan *http.Response, method, path string) {
	defer close(c)
	req, err := http.NewRequest(method, fc.Uri(path), nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := fc.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	c <- res
}

func (fc FacebookClient) Get(path string) chan *http.Response {
	c := make(chan *http.Response)
	go fc.fetchRequest(c, "Get", path)
	return c
}
