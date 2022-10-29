package lib

type Ilib interface {
	Login(username string, passwd string) (bool, error)
}
