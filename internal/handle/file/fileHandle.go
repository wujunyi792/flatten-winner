package file

import (
	"github.com/gin-gonic/gin"
	config "github.com/wujunyi792/flatten-winner/internal/controller/file"
	"github.com/wujunyi792/flatten-winner/internal/response/dto"
	"github.com/wujunyi792/flatten-winner/internal/service/scanFolder"
)

func HandleGetList(c *gin.Context) {
	var res dto.JsonDataResponse
	res.Clear()
	count, list, err := config.GetList()
	if err != nil {
		res.Code = 50001
		res.Message = err.Error()
		c.JSON(res.Code/100, res)
		c.Abort()
		return
	}
	res.Count = int(count)
	res.Data = list
	c.JSON(res.Code/100, res)
}

func HandleRunJob(c *gin.Context) {
	var res dto.JsonResponse
	res.Clear()
	err := scanFolder.LoadAllFolder()
	if err != nil {
		res.Code = 50001
		res.Message = err.Error()
		c.JSON(res.Code/100, res)
		c.Abort()
		return
	}
	c.JSON(res.Code/100, res)
}
