package controller

import (
	mysql "github.com/exiaohao/golang-template/pkg/db"
	"github.com/exiaohao/golang-template/pkg/example/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"time"
)

func Example(c *gin.Context) {
	panic("My panic")
	SuccessResponse(c, "hello, it's default result!")
}


// MyBooks return all my books in my library
func MyBooks(c *gin.Context) {
	mysql.GetDB()
	SuccessResponse(c, 0)
}

func MyBook(c *gin.Context) {
	bookId := c.Param("id")
	var book model.Book
	if err := db.First(&book, bookId).Error; err != nil {
		glog.Infof("Result Error: %s", err)
		FailedResponse(c, err)
		return
	}

	SuccessResponse(c, book)
}

func AddBooks(c *gin.Context) {
	var NewBookForm model.Book

	if err := c.BindJSON(&NewBookForm); err != nil {
		FailedResponse(c, err)
		return
	}
	timeNow := time.Now()

	NewBookForm.CreatedAt = timeNow
	NewBookForm.UpdatedAt = timeNow
	glog.Info(NewBookForm)

	result := db.Save(&NewBookForm)
	if result.Error != nil {
		FailedResponse(c, result.Error)
		return
	}
	SuccessResponse(c, NewBookForm)
}