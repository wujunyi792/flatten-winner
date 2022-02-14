package config

import (
	"github.com/wujunyi792/flatten-winner/internal/db"
	"github.com/wujunyi792/flatten-winner/internal/logger"
	"github.com/wujunyi792/flatten-winner/internal/model/Mysql"
	"gorm.io/gorm"
	"sync"
)

func init() {
	logger.Info.Println("start init Table ...")
	dbManage = GetManage()
}

type DBManage struct {
	mDB     *db.MainGORM
	sDBLock sync.RWMutex
}

var dbManage *DBManage = nil

func (manager *DBManage) getGOrmDB() *gorm.DB {
	return manager.mDB.GetDB().Debug()
}

func (manager *DBManage) atomicDBOperation(op func()) {
	manager.sDBLock.Lock()
	op()
	manager.sDBLock.Unlock()
}

func GetManage() *DBManage {
	if dbManage == nil {
		var Db = db.MustCreateGorm()
		err := Db.GetDB().AutoMigrate(&Mysql.File{})
		if err != nil {
			logger.Error.Fatalln(err)
			return nil
		}
		dbManage = &DBManage{mDB: Db}
	}
	return dbManage
}

func DelAllPathFile(dir string) (err error) {
	err = GetManage().getGOrmDB().Where("base_path = ?", dir).Delete(&Mysql.File{}).Error
	return err
}

func Insert(in *[]Mysql.File) (err error) {
	GetManage().atomicDBOperation(func() {
		err = GetManage().getGOrmDB().CreateInBatches(in, 100).Error
	})
	return err
}

func GetList() (count int64, res []Mysql.File, err error) {
	err = GetManage().getGOrmDB().Model(&Mysql.File{}).Count(&count).Find(&res).Error
	return
}
