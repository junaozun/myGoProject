package user

import (
	"net/http"
	"weblogindemo/errno"
	"weblogindemo/model"
	"weblogindemo/service"

	"github.com/gin-gonic/gin"
)

type Role struct{}

// 用户注册
func (Role) Register(c *gin.Context) {
	data := model.User{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, errno.ERROR.SetMsg("解析错误"))
		return
	}
	status := service.RegisterUser(data)
	c.JSON(http.StatusOK, status)
}

// 用户登录
func (Role) Login(c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, errno.ERROR.SetMsg("解析错误"))
		return
	}
	status, token := service.LoginUser(data)
	c.JSON(http.StatusOK, status.WithData(token))
}
