package shared

import "github.com/gin-gonic/gin"

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
	StatusError   = "error"
)

type Response struct {
	Data    interface{} `json:"data"`
	Status  string   `json:"status"`
	Message string   `json:"message"`
}

type Context struct {
	ResContext *gin.Context
}

func (res *Context) ResponseData(status string, message string, data interface{}) Response {

	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return response
}
