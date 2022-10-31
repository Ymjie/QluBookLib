package job

import (
	"cklib/internal/user"
	"cklib/pkg/logger"
	"context"
	"fmt"
	"sync"
	"time"
)

var Mlog = logger.New(nil, logger.LDEBUG, 0)

func NewJob(u *user.User) func() {
	return func() {
		gog(u)
	}
}

func gog(us *user.User) {
	//us := user.NewUser(username, passwd, list)
	if !us.Login() {
		Mlog.PF(logger.LERROR, "%s", "登陆失败")
		return
	}
	Mlog.PF(logger.LINFO, "%s", "登陆成功")
	timetk := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timetk.C:
			time1, _ := time.ParseInLocation("15:04:05", "6:00:00", time.Local)
			time2, _ := time.ParseInLocation("15:04:05", time.Now().Format("15:04:05"), time.Local)
			//sub := time1.Sub(time2).Seconds()
			//Mlog.PF(logger.LDEBUG, "\r距离Book时间%v", time1.Sub(time2))
			fmt.Printf("\r距离Book时间%v", time1.Sub(time2))
			if time2.Hour() == 6 {
				timetk.Stop()
				go6(us)
				Mlog.PF(logger.LINFO, "程序今天的生命周期已完成，此定时任务退出")
			}
		}
	}

}

func go6(u *user.User) {
	idchan := make(chan int, 10)
	go func(u *user.User) {
		for _, i := range u.BookList {
			idchan <- i
		}
		close(idchan)
	}(u)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go book6(u, idchan, &wg, ctx, cancel)
	}
	wg.Wait()
}

func book6(u *user.User, bidchan chan int, wg *sync.WaitGroup, ctx context.Context, cal context.CancelFunc) {
	defer wg.Done()
	var bid int
	var ok bool
	bid = u.BookList[0]
	for {
		select {
		default:
			Mlog.PF(logger.LINFO, "开始预约：%d", bid)
			bookresp, err := u.Book(bid, 1)
			if err != nil {
				Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, err.Error())
			}
			if bookresp.Msg == "没有登录或登录已超时" {
				Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, bookresp.Msg)
				cal()
			}
			if bookresp.Msg == "该空间当前状态不可预约" {
				Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, bookresp.Msg)
				bid, ok = <-bidchan
				if !ok {
					cal()
				}
			}
			if bookresp.Msg == "当前用户在该时段已存在预约，不可重复预约" {
				Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, bookresp.Msg)
				cal()
			}
			if bookresp.Status == 1 {
				Mlog.PF(logger.LINFO, "账号：%s,已成功Book:%d,%v", u.Username, bid, bookresp)
				cal()
			}
		case <-ctx.Done():
			return

		}

	}
}
