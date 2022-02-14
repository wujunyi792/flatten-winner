package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/flatten-winner/internal/handle/config"
	"github.com/wujunyi792/flatten-winner/internal/handle/file"
	"github.com/wujunyi792/flatten-winner/internal/response/dto"
)

func MainRouter(e *gin.Engine) {
	e.Any("", func(c *gin.Context) {
		res := dto.JsonResponse{}
		res.Clear()
		res.Data = struct {
			UA   string
			Host string
		}{
			UA:   c.Request.Header.Get("User-Agent"),
			Host: c.Request.Host,
		}
		c.JSON(res.Code/100, res)
	})
	v1 := e.Group("v1")
	v1.GET("/list", file.HandleGetList)
	v1.POST("/set", config.HandleSetPath)
	v1.GET("work", file.HandleRunJob)
}
