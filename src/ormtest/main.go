package main

import (
	"orm/mysql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

type UserInfo struct{
	id 	int
	UserName	string
	PassWord	string
}

func (userInfo *UserInfo)GetAddSql() string{
	return "INSERT INTO ormtestdb.userinfo (id, username, password) VALUES(?,?,?)"
}

func (userInfo *UserInfo)GetAddValues() []interface{}{
	return []interface{}{userInfo.id, userInfo.UserName, userInfo.PassWord}
}

func (userInfo *UserInfo)GetDelSql() string{
	return "DELETE FROM ormtestdb.userinfo WHERE id = ?"
}

func (userInfo *UserInfo)GetDelValues() []interface{}{
	return []interface{}{userInfo.id}
}

func (userInfo *UserInfo)GetSaveSql() string{
	return "UPDATE ormtestdb.userinfo SET username = ?, password = ? WHERE id = ?"
}

func (userInfo *UserInfo)GetSaveValues() []interface{}{
	return []interface{}{userInfo.id, userInfo.UserName, userInfo.PassWord}
}

func (userInfo *UserInfo)GetLoadSql() string{
	return "SELECT username, password FROM ormtestdb.userinfo WHERE id = ?"
}

func (userInfo *UserInfo)GetLoadValues() []interface{}{
	return []interface{}{userInfo.id}
}

func (userInfo *UserInfo)GetLoadAddrs() []interface{}{
	return []interface{}{&userInfo.UserName, &userInfo.PassWord}
}

func main()  {
	orm := new(mysql.MySqlOrm)
	orm.Open("ormtest:ormtest@tcp(www.kongyt.com:3306)/ormtestdb")
	//userInfo := new(UserInfo)
	//userInfo.id = 1000
	//userInfo.UserName = "nihoa"
	//userInfo.PassWord = "password2"
	//orm.Add(userInfo)

	userInfo2 := new(UserInfo)
	userInfo2.id = 1000
	orm.Load(userInfo2)
	fmt.Printf("%v\n", userInfo2)

	orm.Close()

}
