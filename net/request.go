package net

import (
	"net/http"
	"net/url"
)

type Request struct {
	url      *url.URL
	header   http.Header
	_type    string
	params   url.Values
	response *http.Response
	*Client
}

func NewRequest(config IConfig) (req *Request, err error) {
	var u *url.URL
	if u, err = config.Url(); err != nil {
		return
	}
	req = &Request{
		url:    u,
		Client: NewClient(),
	}
	return
}

func (self *Request) Url() *url.URL {
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
