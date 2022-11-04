package cron

import (
	"cklib/pkg/logger"
	"github.com/robfig/cron/v3"
)

func Newcron(spec string, job func(), Mlog *logger.MyLogger) {
	Mlog.PF(logger.LINFO, "将在%s运行", spec)
	c := cron.New(cron.WithSeconds())
	//spec := "30 30 5 * * *"
	_, err := c.AddFunc(spec, job)
	if err != nil {
		panic(err)
	}
	c.Start()
}
