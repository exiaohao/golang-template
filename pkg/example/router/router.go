package router

import (
	"github.com/exiaohao/golang-template/pkg/example/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.NoRoute(controller.Example)

	apiV1 := r.Group("/v1")
	apiV1.GET("/books", controller.MyBooks)
	apiV1.GET("/book/:id", controller.MyBook)
	apiV1.POST("/book", controller.AddBooks)
}