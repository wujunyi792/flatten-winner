package scanFolder

import (
	config2 "github.com/wujunyi792/flatten-winner/internal/controller/config"
	config "github.com/wujunyi792/flatten-winner/internal/controller/file"
	"github.com/wujunyi792/flatten-winner/internal/logger"
	"github.com/wujunyi792/flatten-winner/internal/service/fs"
)

func LoadAllFolder() error {
	path, err := config2.GetConfig("path")
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	err = config.DelAllPathFile(path)
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	pData, err := fs.WalkDir(path)
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	err = config.Insert(pData)
	if err != nil {
		logger.Error.Println(err)
		return err
	}
	return nil

}
