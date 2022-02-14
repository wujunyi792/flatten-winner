package fs

import (
	"errors"
	"github.com/wujunyi792/flatten-winner/internal/logger"
	"github.com/wujunyi792/flatten-winner/internal/model/Mysql"
	"os"
	"path/filepath"
)

type Walker struct {
	files []string
}

func WalkDir(dir string) (*[]Mysql.File, error) {
	exists, err := PathExists(dir)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("路径不存在，当前路径: " + GetCurrentAbPath())
	}

	var res []Mysql.File

	walker := new(Walker)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			absPath, _ := filepath.Abs(path)
			walker.files = append(walker.files, path)
			res = append(res, Mysql.File{
				BasePath:       dir,
				FileName:       info.Name(),
				AbsPath:        absPath,
				Size:           info.Size(),
				Mode:           info.Mode().String(),
				LastModifyTime: info.ModTime(),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	logger.Info.Printf("找到 %d 个文件\n", len(walker.files))
	return &res, nil
}
