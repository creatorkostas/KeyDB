package api

import (
	"strconv"
	"strings"

	internal "github.com/creatorkostas/KeyDB/database/database_core/conf"
	"github.com/gin-gonic/gin"
)

type ActionResponce struct {
	Error       error
	Code        int
	From        string
	Description any
}

func (err *ActionResponce) ToString() string {
	var err_string strings.Builder
	err_string.WriteString(strconv.Itoa(err.Code))
	err_string.WriteString(" | ")
	err_string.WriteString(err.Error.Error())
	err_string.WriteString(" | ")
	err_string.WriteString(err.From)
	err_string.WriteString(" | ")
	err_string.WriteString(err.Description.(string))
	err_string.WriteString("\n")
	return err_string.String()
}

type JsonResponce struct {
	Message any `json:"response"`
}

type HttpResponce struct {
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

func (res HttpResponce) SendResponce() {
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
