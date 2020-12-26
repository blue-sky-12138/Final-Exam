package cmd

import (
	"FinalHomework/question3/model"
	"github.com/gin-gonic/gin"
)

var(
	router   *gin.Engine
)

func Entrance(){
	//http://localhost:8080/index
	//预载服务器
	router = gin.Default()
	router.LoadHTMLGlob("question3/templates/html/*")
	router.Static("/css", "./templates/css")

	router.GET("/index",model.Index)

	router.GET("/Login",model.LoginGet)
	router.POST("/Login/Create",model.LoginPost)
	router.GET("/Logout",model.LoginOut)

	router.GET("/Register",model.RegisterGet)
	router.POST("/Register/Create",model.RegisterPost)

	router.GET("/Recharge",model.Recharge)

	router.GET("/History",model.History)

	router.GET("/Deal",model.Deal)
	router.POST("/Deal",model.DealCreate)


	router.GET("/Past",model.Past)

	//运行服务器
	router.Run(":8080")
}
