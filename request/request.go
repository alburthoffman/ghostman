package request

import (
	"fmt"
	"github.com/newly12/ghostman/inventory"
	"github.com/newly12/ghostman/template"
)

type Request struct {
	Host string
	template.Template
}

type BaseResponse struct {
	StatusCode int
	Body       string
}

func (req Request) Geturi() string {
	return fmt.Sprintf("%s://%s:%d%s", req.Protc, req.Host, req.Port, req.Uri)
}

func GetRequests(invt inventory.Inventory, tmpl *template.Template) []Request {
	var requests []Request
	for _, h := range invt.GetAll() {
		requests = append(requests, Request{h, *tmpl})
	}
	return requests
}
