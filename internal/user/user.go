package user

import (
	"cklib/internal/lib"
	"cklib/internal/lib/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Username string
	Passwd   string
	BookList []int
	Lib      *lib.Lib
}

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

func (u *User) Book(id, advanceTime int) (model.Bookresp, error) {
	userid, _ := strconv.Atoi(u.Username)
	bookresp, err := u.Lib.Book(userid, id, advanceTime)
	if err != nil {
		if strings.Contains(err.Error(), "invalid character") {
			err = errors.New("预约系统Web服务出错，未返回正确信息")
		} else if strings.Contains(err.Error(), "closed by the remote host") {
			err = errors.New("预约系统Web服务关闭Connection")
		} else if strings.Contains(err.Error(), "EOF") {
			err = errors.New("预约系统Web服务未响应")
		}
	}

	return bookresp, err

}
func (u *User) GetBooklist() (model.Booklist, error) {
	return u.Lib.GetBooklist()
}
