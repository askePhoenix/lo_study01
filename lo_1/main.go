package main

import (
	"fmt"
	"github.com/samber/lo"
)

type User struct {
	Id int64
	Name string
}

func main() {
	var userArr []User = []User{
		{1, "teddy"},
		{2, "aske"},
		{3, "teddyChang"},
		{4, "teddyChang"},
		{4, "teddyChang"},
		{4, "teddyChang2"},
	}
	names := lo.UniqBy[User](userArr, func(u User) int64 {
		return u.Id
	})

	fmt.Println(names)
}
