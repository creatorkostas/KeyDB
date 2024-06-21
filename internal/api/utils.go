package api

import (
	"github.com/creatorkostas/KeyDB/internal"
	"github.com/gin-gonic/gin"
)

type JsonResponce struct {
	Message any `json:"response"`
}

type Responce struct {
	ErrorMessage string
	Result       any
	C            *gin.Context
	OkCode       int
	ErrorCode    int
	Result_error error
}

// func sendSimpeResponce(c *gin.Context, result any, error_message string) {
// 	if result == nil {
// 		c.JSON(http.StatusBadRequest, JsonResponce{Message: error_message})
// 	} else {
// 		c.JSON(http.StatusOK, result)
// 	}
// 	c.Request.Context().Done()
// }

func (res *Responce) sendResponce() {
	if internal.Send_all_errors_in_requests && res.Result_error != nil {
		res.C.JSON(res.ErrorCode, JsonResponce{Message: res.Result_error.Error()})
	} else {
		if res.Result == nil || res.Result_error != nil {
			res.C.JSON(res.ErrorCode, JsonResponce{Message: res.ErrorMessage})
		} else {
			res.C.JSON(res.OkCode, JsonResponce{Message: res.Result})
		}
	}

	res.C.Request.Context().Done()
}
