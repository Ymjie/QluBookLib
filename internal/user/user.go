package user

import (
	"cklib/internal/lib"
	"cklib/internal/lib/model"
	"errors"
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
	login, _ := u.Lib.Login(u.Username, u.Passwd)
	if !login {
		//fmt.Println(err)
		return false
	}
	go u.keepLive()
	return login
}

func (u *User) Book(id, advanceTime int) (model.Bookresp, error) {
	userid, _ := strconv.Atoi(u.Username)
	bookresp, err := u.Lib.Book(userid, id, advanceTime)
	err = vkhttperr(err)
	return bookresp, err

}
func (u *User) GetBooklist() (model.Booklist, error) {
	booklist, err := u.Lib.GetBooklist()
	err = vkhttperr(err)
	return booklist, err
}
func (u *User) LookBook(id int) (model.Lookbook, error) {
	return u.Lib.LookBook(id)

}

func vkhttperr(err error) error {
	if err != nil {
		if strings.Contains(err.Error(), "invalid character") {
			err = errors.New("解析出错，未返回正确信息")
		} else if strings.Contains(err.Error(), "closed by the remote host") {
			err = errors.New("请求出错，Remote Host关闭Connection")
		} else if strings.Contains(err.Error(), "EOF") {
			err = errors.New("请求出错,未响应")
		}
	}
	return err
}
