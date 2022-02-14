package config

import (
	"github.com/gin-gonic/gin"
	config2 "github.com/wujunyi792/flatten-winner/config"
	"github.com/wujunyi792/flatten-winner/internal/controller/config"
	"github.com/wujunyi792/flatten-winner/internal/response/dto"
	"github.com/wujunyi792/flatten-winner/internal/service/fs"
	"github.com/wujunyi792/flatten-winner/internal/service/scanFolder"
)

func HandleSetPath(c *gin.Context) {
	var res dto.JsonResponse
	res.Clear()
	var req dto.SetPathConfig
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Code = 40200
		res.Message = err.Error()
		c.JSON(res.Code/100, res)
		c.Abort()
		return
	}
	if req.Key != config2.GetConfig().SERVICE.PASSKEY {
		res.Code = 20201
		res.Message = "密钥错误"
		c.JSON(res.Code/100, res)
		c.Abort()
		return
	}
	exist, err := fs.PathExists(req.Path)
	if err != nil {
		res.Code = 20201
		res.Message = err.Error()
		c.JSON(res.Code/100, res)
		c.Abort()
		return
	}
	if !exist {
		res.Message = "路径不存在，当前路径为：" + fs.GetCurrentAbPath()
		res.Code = 20201
		c.JSON(res.Code/100, res)
		c.Abort()
		return
	}

	err = config.SetFilePath(req.Path)
	if err != nil {
		res.Code = 50001
		res.Message = err.Error()
		c.JSON(res.Code/100, res)
		c.Abort()
		return
	}
	_ = scanFolder.LoadAllFolder()
	c.JSON(res.Code/100, res)
}
