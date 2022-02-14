package Mysql

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	gorm.Model
	ConfigName string `gorm:"unique"`
	Value      string
}

func (Config) TableName() string {
	return "configs"
}

type File struct {
	ID        string         `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	BasePath       string    `json:"basePath"`
	AbsPath        string    `json:"absPath"`
	FileName       string    `json:"fileName"`
	Size           int64     `json:"size"`
	Mode           string    `json:"mode"`
	LastModifyTime time.Time `json:"lastModifyTime"`
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.NewV4().String()
	return
}
