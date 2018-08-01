package execute

import (
	"bytes"
	"github.com/newly12/ghostman/logging"
	"github.com/newly12/ghostman/request"
	"github.com/newly12/ghostman/response"
	"github.com/newly12/ghostman/utils"
	"io/ioutil"
	"net/http"
)

func Execute(req *request.Request) *response.Response {
	log := logging.GetLogger()
	log.Debugf("working on %s", req.Geturi())

	res := response.Response{}
	res.Request = *req
	res.StatusCode = -1

	bodybyte := []byte(req.Body)
	httpreq, err := http.NewRequest(req.Method, req.Geturi(), bytes.NewBuffer(bodybyte))

	// set Header
	for _, headermap := range req.Headers {
		for k, v := range headermap {
			httpreq.Header.Set(k, v)
		}
	}
	// if req.CT != "" {
	// 	httpreq.Header.Set("Content-Type", req.CT)
	// }
	// if req.Auth != "" {
	// 	httpreq.Header.Set("Authorization", req.Auth)
	// }

	resp, err := utils.HttpClient.Do(httpreq)

	if err != nil {
		log.Errorf("Got error: %s", err)
		res.Body = err.Error()
		return &res
	}
	res.StatusCode = resp.StatusCode
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			log.Error(err2)
		} else {
			res.ResponseBody = string(bodyBytes)
		}
	} else {
		res.ResponseBody = "na"
	}

	return &res
}
