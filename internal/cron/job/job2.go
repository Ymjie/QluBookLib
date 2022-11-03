package job

import (
	"cklib/internal/user"
	"cklib/pkg/logger"
	"cklib/pkg/notice"
	"context"
	"fmt"
	"sync"
	"time"
)

func NewJob(user *user.User, nt notice.Notifier, t int, myLogger *logger.MyLogger) *Job {
	one := &Job{
		Nt:   nt,
		Mlog: myLogger,
		u:    user,
		t:    t,
	}
	myLogger.PF(logger.LINFO, "将使用%d线程并发 预约List:%v", t, user.BookList)
	return one
}

func (j *Job) Start() {
	if !j.u.Login() {
		j.Mlog.PF(logger.LERROR, "%s", "登陆失败")
		return
	}
	j.Mlog.PF(logger.LINFO, "%s", "登陆成功")
	j.gog()
}

func (j *Job) gog() {
	timetk := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timetk.C:
			//time1, _ := time.ParseInLocation("15:04:05", "6:00:00", time.Local)
			time2, _ := time.ParseInLocation("15:04:05", time.Now().Format("15:04:05"), time.Local)
			//sub := time1.Sub(time2).Seconds()
			//Mlog.PF(logger.LDEBUG, "\r距离Book时间%v", time1.Sub(time2))
			//fmt.Printf("\r距离Book时间%v", time1.Sub(time2))
			if time2.Hour() == 6 {
				timetk.Stop()
				j.go6()
				var isok bool
				isok = true
				for isok {
					msg, b := ckBook(j.u)
					if b {
						isok = false
						j.Mlog.PF(logger.LINFO, "%s", msg)
						if j.Nt.Getenable() {
							j.Mlog.PF(logger.LINFO, "发送通知")
							j.Nt.SendNotify(fmt.Sprintf("[%s]%s", j.u.Username, msg))
						}

					}
				}

				j.Mlog.PF(logger.LINFO, "程序今天的生命周期已完成，此定时任务退出")
			}
		}
	}
}

func (j *Job) go6() {
	idchan := make(chan int, 10)
	go func(u *user.User) {
		for _, i := range u.BookList {
			idchan <- i
		}
		close(idchan)
	}(j.u)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < j.t; i++ {
		wg.Add(1)
		go j.book6(idchan, &wg, ctx, cancel)
	}
	wg.Wait()
}

func (j *Job) book6(bidchan chan int, wg *sync.WaitGroup, ctx context.Context, cal context.CancelFunc) {
	defer wg.Done()
	var bid int
	var ok bool
	bid = j.u.BookList[0]
	for {
		select {
		default:
			j.Mlog.PF(logger.LINFO, "开始预约：%d", bid)
			bookresp, err := j.u.Book(bid, 1)
			//fmt.Println(bookresp)
			if err != nil {
				j.Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, err.Error())
			}
			if bookresp.Msg == "没有登录或登录已超时" {
				j.Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, bookresp.Msg)
				cal()
			}
			if bookresp.Msg == "该空间当前状态不可预约" {
				j.Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, bookresp.Msg)
				bid, ok = <-bidchan
				if !ok {
					cal()
				}
			}
			if bookresp.Msg == "当前用户在该时段已存在预约，不可重复预约" {
				j.Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, bookresp.Msg)
				cal()
			}
			if bookresp.Msg == "预约时间段不存在！" {
				j.Mlog.PF(logger.LINFO, "预约：%d 失败,%s", bid, bookresp.Msg)
				cal()
			}
			if bookresp.Status == 1 {
				j.Mlog.PF(logger.LINFO, "账号：%s,已成功Book:%d,%v", j.u.Username, bid, bookresp)
				cal()
			}
		case <-ctx.Done():
			return

		}

	}
}

func ckBook(u *user.User) (string, bool) {
	booklist, err := u.GetBooklist()
	if err != nil {
		return err.Error(), false
	}
	if len(booklist) == 0 {
		return "null", false
	}
	if booklist[0].Status == "预约成功" {
		return fmt.Sprintf("成功预约：%s", booklist[0].Area), true
	}
	return "预约失败咯~www~", true
}
