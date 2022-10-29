package user

import (
	"cklib/internal/lib"
	"cklib/pkg/logger"
	"errors"
	"fmt"
	"strconv"
)

type User struct {
	Username string
	Passwd   string
	BookList []int
	Lib      *lib.Lib
}

var Mlog = logger.New(nil, logger.LDEBUG, 0)

func NewUser(username, passwd string, bid []int) *User {
	varlib := lib.NewLib()
	return &User{
		Username: username,
		Passwd:   passwd,
		BookList: bid,
		Lib:      varlib,
	}

}

func (u *User) keepLive() {

}

func (u *User) Login() bool {
	login, err := u.Lib.Login(u.Username, u.Passwd)
	if !login {
		fmt.Println(err)
		return false
	}
	go u.keepLive()
	return login
}

func (u *User) Book(id, advanceTime int) (bool, error) {
	userid, _ := strconv.Atoi(u.Username)
	bookresp, err := u.Lib.Book(userid, id, advanceTime)
	if err != nil {
		return false, err
	}
	if bookresp.Msg == "该空间当前状态不可预约" {
		return false, errors.New("1")
	}
	if bookresp.Msg == "没有登录或登录已超时" {
		return false, errors.New("2")
	}

	if bookresp.Status != 1 {
		return false, errors.New(fmt.Sprintf("%v", bookresp))
	}

	return true, nil
}
