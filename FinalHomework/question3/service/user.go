package service

import (
	"FinalHomework/question3/database"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type User struct {					//用户信息结构体
	ID int
	Name string
	Password string
}
type DataCryptographyMD5 struct {	//MD5加密结构体
	Data string
	Result string
}
var (
	Dc     DataCryptographyMD5
	tem database.Tem
)

//注册
func Register(context *gin.Context)string{
	registerTelephoneNumber,_:= strconv.Atoi(context.PostForm("telephone_number"))
	registerName:=context.PostForm("name")
	registerPassword:=context.PostForm("password")
	registerPasswordAgain:=context.PostForm("passwordAgain")

	//简易判断手机号是否被注册
	name:= database.FindUserName(registerTelephoneNumber)
	if name!=""{
		return "该手机号已注册"
	}

	//简易检查密码
	if strings.Contains(registerPassword,"/"){
		return "密码含有非法字符"
	}else if len(registerPassword)<=6{
		return "密码过短"
	}
	if registerPasswordAgain!=registerPassword{
		return "两次密码不一致"
	}

	Dc.Cryptography()
	database.InsertUser(registerName,Dc.Result,registerTelephoneNumber)
	return ""
}

//检查用户是否存在
func (user *User)CheckUser()bool{
	OKstring:= database.FindUserName(user.ID)
	if OKstring!=""{
		user.Name=OKstring
		return true
	}
	return false
}

//检查密码是否正确
func (user *User)CheckPassword()bool{
	tem.TemString= database.FindUserPassword(user.Name)
	if Dc.Cryptography(); tem.TemString== Dc.Result{
		return true
	}
	return false
}