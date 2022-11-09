package model

import (
	"time"
)

type Lookbook struct {
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
