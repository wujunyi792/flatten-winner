package config

import (
	"errors"
	"github.com/wujunyi792/flatten-winner/internal/db"
	"github.com/wujunyi792/flatten-winner/internal/logger"
	"github.com/wujunyi792/flatten-winner/internal/model/Mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return manager.mDB.GetDB()
}

func (manager *DBManage) atomicDBOperation(op func()) {
	manager.sDBLock.Lock()
	op()
	manager.sDBLock.Unlock()
}

func GetManage() *DBManage {
	if dbManage == nil {
		var Db = db.MustCreateGorm()
		err := Db.GetDB().AutoMigrate(&Mysql.Config{})
		if err != nil {
			logger.Error.Fatalln(err)
			return nil
		}
		dbManage = &DBManage{mDB: Db}
	}
	return dbManage
}

func SetFilePath(path string) (err error) {
	err = GetManage().getGOrmDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "config_name"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
	}).Create(&Mysql.Config{
		ConfigName: "path",
		Value:      path,
	}).Error
	return
}

func SetConfig(key string, value string) (err error) {
	err = GetManage().getGOrmDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "config_name"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "updated_at"}),
	}).Create(&Mysql.Config{
		ConfigName: key,
		Value:      value,
	}).Error
	return
}

func GetConfig(key string) (value string, err error) {
	var v Mysql.Config
	e := GetManage().getGOrmDB().Model(&Mysql.Config{}).Where("config_name = ?", key).Find(&v)
	if v.ID == 0 {
		err = errors.New(key + " 记录不存在")
		return
	}
	return v.Value, e.Error
}
