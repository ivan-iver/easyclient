package net

import (
	"log"
	"net/http"
	"net/url"
)

type Request struct {
	url      string
	header   http.Header
	_type    string
	params   url.Values
	response *http.Response
	*Client
}

func NewRequest(config IConfig) *Request {
	log.Printf("Url: %v", config.Url())
	return &Request{
		url:    config.Url(),
		Client: NewClient(),
	}
}

func (self *Request) Url() string {
	return self.url
}

func (self *Request) Header() http.Header {
	return self.header
}

func (self *Request) Type() string {
	return self._type
}

func (self *Request) Params() url.Values {
	return self.params
}

func (self *Request) Response() *http.Response {
	return self.response
}

func (self *Request) Send() (err error) {
	self.response, err = self.Post(self)
	return
}
