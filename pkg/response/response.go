package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseOKWithDataModel struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseOkModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseErrorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseErrorCustomModel struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ResponseOkWithData(c *gin.Context, data interface{}) {
	response := ResponseOKWithDataModel{
		Code:    200,
		Data:    data,
		Message: "OK",
	}
	c.JSON(http.StatusCreated, response)
}

func ResponseCreated(c *gin.Context, data interface{}) {
	response := ResponseOKWithDataModel{
		Code:    201,
		Data:    data,
		Message: "Created",
	}
	c.JSON(http.StatusCreated, response)
}

func ResponseOK(c *gin.Context, message string) {
	response := ResponseOkModel{
		Code:    200,
		Message: message,
	}
	c.JSON(http.StatusOK, response)
}

func ResponseError(c *gin.Context, err string, code int) {
	response := ResponseErrorModel{
		Code:    500,
		Message: err,
	}
	c.JSON(code, response)
}

func ResponseCustomError(c *gin.Context, err interface{}, code int) {
	response := ResponseErrorCustomModel{
		Code:    500,
		Message: err,
	}
	c.JSON(code, response)
}
