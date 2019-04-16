package controller

import (
	mysql "github.com/exiaohao/golang-template/pkg/db"
	"github.com/exiaohao/golang-template/pkg/example/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Example(c *gin.Context) {
	SuccessResponse(c, nil)
}


// MyBooks return all my books in my library
func MyBooks(c *gin.Context) {
	mysql.GetDB()
	SuccessResponse(c, 0)
}

func AddBooks(c *gin.Context) {
	glog.Info("Add books called!")
	var newBook model.Book
	if err := c.BindJSON(newBook); err != nil {
		glog.Info(err)
	} else {
		glog.Info(newBook)
	}
	SuccessResponse(c, newBook)
}