package helpers

import (
	modelRef "go-rest-starter/src/api/models/req/auth"

	"github.com/gin-gonic/gin"
)

func ParseExtraClientInfo(req *gin.Context, d *modelRef.AuthInitReq) {
	d.IP = req.RemoteIP()

	d.UserAgent = req.GetHeader("User-Agent")
	d.Referer = req.GetHeader("Referer")

	if len(req.GetHeader("Host")) > 0 {
		d.Host = req.GetHeader("Host")
	} else {
		d.Host = req.GetHeader("Origin")
	}

	// for k, vals := range req.Request.Header {
	// 	if strings.Contains(k, "Origin") {
	// 		fmt.Printf("%s", k)
	// 		// fmt.Printf(" %v", strings.Contains(k, "Origin"))
	// 		for _, v := range vals {
	// 			fmt.Printf("\t%s\n", v)
	// 		}
	// 	}
	// }
}
