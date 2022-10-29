package cron

import (
	"cklib/pkg/logger"
	"github.com/robfig/cron/v3"
)

var Mlog = logger.New(nil, logger.LDEBUG, 0)

func Newcron(spec string, job func()) {
	c := cron.New(cron.WithSeconds())
	//spec := "30 30 5 * * *"
	c.AddFunc(spec, job)
	Mlog.PF(logger.LINFO, "将在%s开始运行", spec)
	c.Start()
}
