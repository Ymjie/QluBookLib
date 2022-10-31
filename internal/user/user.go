package user

import (
	"cklib/internal/lib"
	"cklib/internal/lib/model"
	"cklib/pkg/logger"
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

func (u *User) Book(id, advanceTime int) (model.Bookresp, error) {
	userid, _ := strconv.Atoi(u.Username)
	bookresp, err := u.Lib.Book(userid, id, advanceTime)
	return bookresp, err

}
