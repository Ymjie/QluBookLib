package http

import (
	"net/http"
	"net/url"
	"strings"
)

type HttpNt struct {
	Enable  bool
	UrlText string
	Mark    string
}

func NewNt(url, mark string) *HttpNt {
	enable := false
	if url != "" {
		enable = true
	}

	return &HttpNt{
		Enable:  enable,
		UrlText: url,
		Mark:    mark,
	}
}

func (h *HttpNt) SendNotify(msg string) {
	if h.Enable {
		msg = url.QueryEscape(msg)
		replace := strings.Replace(h.UrlText, h.Mark, msg, -1)
		http.Get(replace)
	}
}
func (h *HttpNt) Getenable() bool {
	return h.Enable
}
