package net

import (
	"log"
	"net/http"
	"strings"
)

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{
		Client: &http.Client{},
	}
}

func (self *Client) Post(request IRequest) (response *http.Response, err error) {
	var req *http.Request
	body := strings.NewReader(request.Params().Encode())
	url := request.Url()
	log.Printf("Body: %v", body)
	if req, err = http.NewRequest("POST", url, body); err != nil {
		return
	}
	req.Header = request.Header()
	return self.Do(req)
}
