package models

import (
"time"
"fmt"
)

type Session struct {
	Id int
	Time time.Time
	ScreenId int
}


func SomeMethod() {
	fmt.Println("hello")
}
