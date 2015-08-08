package net

import (
	"net/http"
	"net/url"
)

type IConfig interface {
	Url() (*url.URL, error)
	Api() string
	User() string
	Pwd() string
	Signature() string
}

type IRequest interface {
	Url() *url.URL
	Header() http.Header
	Type() string
	Params() url.Values
	Response() *http.Response
	Send() error
}
