package service

import (
	"FinalHomework/question3/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func FindHistory(context *gin.Context)*database.Histories{
	var histories database.Histories
	name,_:=context.Cookie("users")
	database.FindAllHistory(&histories,name)
	return &histories
}

func Deal(context *gin.Context)string{
	sender,_:=context.Cookie("users")
	receiver:=context.PostForm("Receiver")
	value, _ :=strconv.Atoi(context.PostForm("Value"))
	remark:=context.PostForm("Remark")

	NowTime:=time.Now()
	time:=fmt.Sprintf("%d-%d-%d %d:%d:%d",NowTime.Year(),NowTime.Month(),NowTime.Day(),NowTime.Hour(),NowTime.Minute(),NowTime.Second())
	database.InsertHistory(sender,time,receiver,value,remark)
	return ""
}