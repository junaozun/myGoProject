package routes

import (
	ctl_user "weblogindemo/controller/user"
	"weblogindemo/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r1 := r.Group("/user")
	roleCtl := ctl_user.Role{}
	{
		r1.POST("register", roleCtl.Register)
		r1.POST("login", roleCtl.Login)
	}

	//r1.Use(middleware.JwtToken())
	/*
	  需要使用中间件的路由
	*/
	return r
}
