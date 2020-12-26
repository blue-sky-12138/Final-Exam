package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Tem struct {
	TemString string
	TemInt    int
	TemSalt   int64
	TemBool bool
}

var (
	tem Tem
)

//寻找用户名
func FindUserName(ID int)string{
	prepare:=fmt.Sprintf("select name from users_information where telephone_number=%d",ID)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindUserID error",err)
	defer stmt.Close()
	if stmt.Next(){
		//存储用户名
		stmt.Scan(&tem.TemString)
		return tem.TemString
	}else{
		return ""
	}
}

//寻找用户手机号
func FindUserTelephone(name string)int{
	prepare:=fmt.Sprintf("select telephone_number from users_information where name='%s'",name)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindUserID error",err)
	defer stmt.Close()
	if stmt.Next(){
		//存储用户名
		stmt.Scan(&tem.TemInt)
		return tem.TemInt
	}else{
		return 0
	}
}

//寻找用户密码
func FindUserPassword(name string)string{
	prepare:=fmt.Sprintf("select password from users_information where name='%s'",name)
	stmt,err:= DataBase.Query(prepare)
	CheckError("FindUserPassword error",err)
	defer stmt.Close()
	if stmt.Next(){
		stmt.Scan(&tem.TemString)
	}
	return tem.TemString
}

//添加注册数据
func InsertUser(name string,password string,telephoneNumber int){
	prepare:=fmt.Sprintf("insert users_information (name,password,telephone_number)value('%s','%s',%d)",name,password,telephoneNumber)
	stmt,err:= DataBase.Prepare(prepare)
	defer stmt.Close()
	CheckError("InsertUser error",err)
	stmt.Exec()
}