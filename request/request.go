package request

import (
	// "bytes"
	"fmt"
	"github.com/newly12/ghostman/inventory"
	"github.com/newly12/ghostman/template"
	// "github.com/newly12/ghostman/utils"
	// "io/ioutil"
	// "log"
	// "net/http"
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
	return fmt.Sprintf("%s://%s:%d%s", req.Protc, req.Host, req.Port, req.Url)
}

// func (req *Request) run() *BaseResponse {
// 	// log.Printf("working on %s", req.Host)
// 	baseres := BaseResponse{}
// 	baseres.StatusCode = -1

// 	bodybyte := []byte(req.Body)
// 	httpreq, err := http.NewRequest(req.Method, req.geturi(), bytes.NewBuffer(bodybyte))

// 	// set Header
// 	if req.CT != "" {
// 		httpreq.Header.Set("Content-Type", req.CT)
// 	}
// 	if req.Auth != "" {
// 		httpreq.Header.Set("Authorization", req.Auth)
// 	}

// 	resp, err := utils.HttpClient.Do(httpreq)

// 	if err != nil {
// 		// log.Printf("got error: %s", err)
// 		baseres.Body = err.Error()
// 		return &baseres
// 	}
// 	baseres.StatusCode = resp.StatusCode
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
// 		if err2 != nil {
// 			log.Println(err2)
// 		} else {
// 			baseres.Body = string(bodyBytes)
// 		}
// 	} else {
// 		baseres.Body = "na"
// 	}

// 	return &baseres
// }

// type requestable interface {
// 	geturi() string
// 	run() *BaseResponse
// }

func GetRequests(invt inventory.Inventory, tmpl *template.Template) []Request {
	var requests []Request
	for _, h := range invt.GetAll() {
		requests = append(requests, Request{h, *tmpl})
	}
	return requests
}

// func GenRequests(t *template.Template, hs Hosts) []Request {
// 	var requests []Request
// 	for _, h := range hs {
// 		requests = append(requests, Request{h, *t})
// 	}
// 	return requests
// }

// func Run(req requestable) *BaseResponse {
// 	return req.run()
// }
