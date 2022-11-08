package main

import (
	"cklib/internal/user"
	"fmt"
)

func main() {
	newUser := user.NewUser("0", "0", []int{})
	newUser.Login()
	timeidf := 2841082
	for {
		book, _ := newUser.LookBook(timeidf)
		timeidf++
		fmt.Printf("%v", book)
	}
}
