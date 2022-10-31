package main

import (
	"cklib/internal/cron"
	"cklib/internal/cron/job"
	"cklib/internal/user"
	"cklib/pkg/logger"
	"os"
	"strconv"
)

func main() {
	var lists []int
	var username string
	var passwd string
	if len(os.Args) > 3 {
		username = os.Args[1]
		passwd = os.Args[2]

		for i := 3; i < len(os.Args); i++ {
			atoi, err := strconv.Atoi(os.Args[i])
			if err != nil {
				return
			}
			lists = append(lists, atoi)

		}
	} else {
		return
	}
	//检查账号密码是否可以正常登陆
	f := user.NewUser(username, passwd, lists)
	if !f.Login() {
		return
	}
	logger.New(nil, logger.LDEBUG, 0).PF(logger.LINFO, "账号: U:%s,P:%s    预约List:%v", username, passwd, lists)

	job1 := job.NewJob(f)
	//job1()
	cron.Newcron("30 30 5 * * *", job1)
	select {}
}
