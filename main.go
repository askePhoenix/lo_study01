package main

import (
	"fmt"
	"github.com/samber/lo"
)

type user struct {
	id int64
	name string
}

type userArray []*user
type mapIdName map[int64]string

func (u userArray) Associate(f func(*user) (int64, string)) mapIdName {
	return lo.Associate[*user, int64, string](u, f)
}

func getIdAndName(u *user) (int64, string) {
	return u.id, u.name
}

func main() {
	var u userArray = []*user{{name: "aske", id: 1}, {name: "teddy", id: 2}}

	userMap := u.Associate(getIdAndName)

	fmt.Println(userMap)
}
