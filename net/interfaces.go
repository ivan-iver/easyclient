package net

import (
	"net/http"
	"net/url"
)

type IConfig interface {
	Url() string
	Api() string
	User() string
	Pwd() string
	Signature() string
}

type IRequest interface {
	Url() string
	Header() http.Header
	Type() string
	Params() url.Values
	Response() *http.Response
	Send() error
}
