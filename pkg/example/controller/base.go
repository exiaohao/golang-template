package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}


func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code: 	1,
		Status: "ok",
		Data:	data,
	})
}