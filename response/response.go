package response

import (
	"github.com/newly12/ghostman/request"
)

type Response struct {
	request.Request
	StatusCode   int
	ResponseBody string
}
