package router

import (
	"github.com/exiaohao/golang-template/pkg/example/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	r.NoRoute(controller.Example)
}