package model

import (
	"FinalHomework/question3/database"
	"FinalHomework/question3/service"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	login       service.User
	information   struct{			//主页template信息结构体
		Title string
		VisterOK bool
		LoginOK bool
	}
	PastContent struct{Content string}	//过度template信息结构体
)

//主页
func Index(context *gin.Context) {
	tem, err :=context.Request.Cookie("users")
	if err==nil{
		if login.ID==0{
			//获取cookie中的用户数据
			login.Name=tem.Value
			login.ID= database.FindUserTelephone(login.Name)
		}
		information.Title="欢迎回来"+login.Name
		information.LoginOK=true
		information.VisterOK=false
		context.HTML(http.StatusOK,"index.html",information)
	}else{
		information.Title="你好游客"
		information.LoginOK=false
		information.VisterOK=true
		context.HTML(http.StatusOK,"index.html",information)
	}
}

func LoginOut(context *gin.Context) {
	cookie, _ :=context.Request.Cookie("users")
	//删除cookie
	context.SetCookie(cookie.Name,cookie.Value,-1,cookie.Path,cookie.Domain,cookie.Secure,cookie.HttpOnly)
	PastContent.Content="注销成功"
	context.Redirect(http.StatusMovedPermanently,"/Past")
}

//登录界面
func LoginGet(context *gin.Context) {
	context.HTML(http.StatusOK,"login.html",nil)
}
//接受传入登录数据
func LoginPost(context *gin.Context) {
	temName:=context.PostForm("name")
	login.ID,_= strconv.Atoi(temName)
	login.Password=context.PostForm("password")

	//登录检测并重定向
	if login.CheckUser() && login.CheckPassword(){
		cookie := &http.Cookie{
			Name:     "users",
			Value:    login.Name,
			MaxAge:   100000,
			Path:     "/",
			Domain:   "localhost",
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(context.Writer,cookie)
		context.Redirect(http.StatusMovedPermanently,"/index")
	}else{
		PastContent.Content="请检查你的用户名和密码是否正确"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//注册界面
func RegisterGet(context *gin.Context) {
	context.HTML(http.StatusOK,"register.html",nil)
}
func RegisterPost(context *gin.Context) {
	err:= service.Register(context)
	switch {
	case err=="":
		PastContent.Content="注册成功"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	default:
		PastContent.Content=err
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}

//充值界面
func Recharge(context *gin.Context){
	context.HTML(http.StatusOK,"register.html",nil)
}

//历史记录
func History(context *gin.Context){
	history:=service.FindHistory(context)
	context.HTML(http.StatusOK,"History.html",history)
}

//交易
func Deal(context *gin.Context){
	context.HTML(http.StatusOK,"Deal.html",nil)
}
func DealCreate(context *gin.Context){
	err:= service.Deal(context)
	switch {
	case err=="":
		PastContent.Content="交易成功"
		context.Redirect(http.StatusMovedPermanently,"/Past")
	default:
		PastContent.Content=err
		context.Redirect(http.StatusMovedPermanently,"/Past")
	}
}


//过渡页面(3秒后回到主页)
func Past(context *gin.Context) {
	context.HTML(http.StatusOK,"Past.html",PastContent)
}
