package v1

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/flatten-winner/internal/handle/config"
	"github.com/wujunyi792/flatten-winner/internal/handle/file"
)

func MainRouter(e *gin.Engine) {
	e.Use(static.Serve("/", static.LocalFile("./frontend", false)))

	v1 := e.Group("v1")
	v1.GET("/list", file.HandleGetList)
	v1.POST("/set", config.HandleSetPath)
	v1.GET("/work", file.HandleRunJob)
}
