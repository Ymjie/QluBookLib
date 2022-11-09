package job

import (
	"cklib/internal/user"
	"cklib/pkg/logger"
	"cklib/pkg/notice"
	"context"
	"encoding/json"
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

	if !user.Login() {
		myLogger.PF(logger.LWARN, "%s", "登陆失败！请检查账号密码和网络状态！")
	}
	myLogger.PF(logger.LINFO, "将使用%d线程并发 预约List:%v", t, user.BookList)
	return one
}

func (j *Job) Start() {
	if !j.u.Login() {
		j.Mlog.PF(logger.LWARN, "%s", "登陆失败")
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
				msg := j.ckBook(j.u)
				if j.Nt.Getenable() {
					j.Mlog.PF(logger.LINFO, "发送通知")
					j.Nt.SendNotify(fmt.Sprintf("[%s]%s", j.u.Username, msg))
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
		go j.book6(idchan, &wg, ctx, cancel, i)
	}
	wg.Wait()
}

func (j *Job) book6(bidchan chan int, wg *sync.WaitGroup, ctx context.Context, cal context.CancelFunc, gonum int) {
	defer wg.Done()
	var bid int
	var ok bool
	bid = j.u.BookList[0]
	num := 1
	for {
		select {
		default:
			prefix := fmt.Sprintf("线程%d,第%d次预约:", gonum, num)
			j.Mlog.PF(logger.LINFO, "%s %d 开始！", prefix, bid)
			bookresp, err := j.u.Book(bid, 1)
			if err != nil {
				j.Mlog.PF(logger.LINFO, "%s %d 失败,%s", prefix, bid, err.Error())
			}
			if bookresp.Msg == "没有登录或登录已超时" {
				j.Mlog.PF(logger.LINFO, "%s %d 失败,%s", prefix, bid, bookresp.Msg)
				cal()
			}
			if bookresp.Msg == "该空间当前状态不可预约" {
				j.Mlog.PF(logger.LINFO, "%s %d 失败,%s", prefix, bid, bookresp.Msg)
				bid, ok = <-bidchan
				if !ok {
					cal()
				}
			}
			if bookresp.Msg == "当前用户在该时段已存在预约，不可重复预约" {
				j.Mlog.PF(logger.LINFO, "%s %d 失败,%s", prefix, bid, bookresp.Msg)
				cal()
			}
			if bookresp.Msg == "预约时间段不存在！" {
				j.Mlog.PF(logger.LINFO, "%s %d 失败,%s", prefix, bid, bookresp.Msg)
				cal()
			}
			if bookresp.Status == 1 {
				bytes, _ := json.Marshal(bookresp)
				j.Mlog.PF(logger.LINFO, "%s %d 成功,详情:%v", prefix, bid, string(bytes))
				cal()
			}
			num++
		case <-ctx.Done():
			return

		}

	}
}

func (j *Job) ckBook(u *user.User) string {
	var count int
	for count <= 100 {
		booklist, err := u.GetBooklist()
		if err != nil {
			j.Mlog.PF(logger.LDEBUG, "%s", err.Error())
			count++
			continue
		}
		if booklist[0].Status == "预约成功" {
			return fmt.Sprintf("成功预约：%s", booklist[0].Area)
		} else if booklist[0].Status != "" {
			b, _ := json.Marshal(booklist[0])
			j.Mlog.PF(logger.LDEBUG, "当前最新预约记录：%s", string(b))
			return "预约失败咯~www~"
		}
	}
	return "查询是否预约成功超时，请自行检查"
}
