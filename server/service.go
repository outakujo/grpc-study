package main

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	err := db.AutoMigrate(&SysInfo{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&ScriptResult{})
	if err != nil {
		panic(err)
	}
	return &Service{db: db}
}

func (r *Service) SaveSysInfo(info *SysInfo) error {
	return r.db.Create(info).Error
}

func (r *Service) SaveScriptResult(info *ScriptResult) error {
	return r.db.Create(info).Error
}

var service *Service
