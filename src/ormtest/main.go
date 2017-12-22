package main

import (
	"fmt"
	"orm"
)



type UserInfo struct{
	Id 	int				`field:"id" index:"pk" auto:"true" table:"userinfo"`
	UserName	string	`field:"username"`
	PassWord	string	`field:"password"`
	TestField	string
}

func main()  {
	orm := orm.NewOrm()

	orm.Register(&UserInfo{})

	orm.Open("mysql", "ormtest:ormtest@tcp(www.kongyt.com:3306)/ormtestdb")
	defer orm.Close()

	userInfo1 := new(UserInfo)
	userInfo1.UserName = "kongyt"
	userInfo1.PassWord = "password1"
	orm.Add(userInfo1)

	userInfo2 := new(UserInfo)
	userInfo2.Id = userInfo1.Id
	orm.Load(userInfo2)

	fmt.Println(userInfo1)
	fmt.Println(userInfo2)

	userInfo1.UserName = "newname"
	orm.Save(userInfo1)

	orm.Del(userInfo1)

}
