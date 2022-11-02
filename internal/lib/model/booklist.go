package model

type Booklist []OneBook

type OneBook struct {
	ID        string
	Area      string
	Begintime string
	Endtime   string
	Status    string
}
