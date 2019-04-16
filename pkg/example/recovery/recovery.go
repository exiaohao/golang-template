package recovery

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io"
)

func Recovery(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return RecoveryWithWriter(f, gin.DefaultErrorWriter)
}

func RecoveryWithWriter(f func(c *gin.Context, err interface{}), out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				glog.Errorf("Error occured: %s", err)
				f(c, err)
			}
		}()
		c.Next()
	}

}