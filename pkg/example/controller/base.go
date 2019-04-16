package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	mysql "github.com/exiaohao/golang-template/pkg/db"
)

var db = mysql.GetDB()

type BaseResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}


func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code: 	200,
		Status: "ok",
		Data:	data,
	})
}

func FailedResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusInternalServerError, BaseResponse{
		Code: 	500,
		Status: "failed",
		Data: 	data,
	})
}
