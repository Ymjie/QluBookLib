package library

import (
	"gorm.io/gorm"
	"time"
)

type Lookbook struct {
	gorm.Model
	LBKID       int
	Status      string
	StuID       int
	Name        string
	Area        string
	Type        string
	BeginTime   time.Time
	EndTime     time.Time
	AuditTime   time.Time
	Auditresult string
}
