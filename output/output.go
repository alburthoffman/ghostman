package output

import (
	"github.com/newly12/ghostman/logging"
	"github.com/newly12/ghostman/response"
)

var log = logging.GetLogger()

func Println(res *response.Response) {
	log.Infof("%s - %d %s", res.Geturi(), res.StatusCode, res.ResponseBody)
}
