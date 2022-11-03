package main

import (
	"cklib/internal/config"
	"cklib/internal/cron"
	"cklib/internal/cron/job"
	"cklib/internal/user"
	"cklib/pkg/logger"
	"cklib/pkg/notice/http"
	"cklib/service"
	"fmt"
	"os"
	"strconv"
)

func main() {
	echoBanner()
	var lists []int
	var username string
	var passwd string
	var configPath string
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
		//检查账号密码是否可以正常登陆
		f := user.NewUser(username, passwd, lists)
		if !f.Login() {
			return
		}
		//f.GetBooklist()
		Mlog := logger.New(nil, logger.LDEBUG, 0)
		Mlog.PF(logger.LINFO, "账号: U:%s,P:%s    预约List:%v", username, passwd, lists)
		job1 := job.NewJob(f, http.NewNt("", ""), 5, Mlog)
		cron.Newcron("30 30 5 * * *", job1.Start, Mlog)
		select {}
	} else if len(os.Args) == 2 {
		configPath = os.Args[1]
		config, err := config.Load(configPath)
		if err != nil {
			panic(err)
		}
		s := service.New(config)
		s.Start()
	}

}

func echoBanner() {
	fmt.Printf("\x1bc")
	fmt.Print("   ███████     ██         ██████                     ██     ██       ██ ██     \n  ██░░░░░██   ░██        ░█░░░░██                   ░██    ░██      ░░ ░██     \n ██     ░░██  ░██ ██   ██░█   ░██   ██████   ██████ ░██  ██░██       ██░██     \n░██      ░██  ░██░██  ░██░██████   ██░░░░██ ██░░░░██░██ ██ ░██      ░██░██████ \n░██    ██░██  ░██░██  ░██░█░░░░ ██░██   ░██░██   ░██░████  ░██      ░██░██░░░██\n░░██  ░░ ██   ░██░██  ░██░█    ░██░██   ░██░██   ░██░██░██ ░██      ░██░██  ░██\n ░░███████ ██ ███░░██████░███████ ░░██████ ░░██████ ░██░░██░████████░██░██████ \n  ░░░░░░░ ░░ ░░░  ░░░░░░ ░░░░░░░   ░░░░░░   ░░░░░░  ░░  ░░ ░░░░░░░░ ░░ ░░░░░   \n\n本项目只是本人个人学习开发并维护，本人不保证任何可用性，也不对使用本软件造成的任何后果负责。\n=====================================\n")
}
