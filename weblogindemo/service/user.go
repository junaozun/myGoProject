package service

import (
	"weblogindemo/encry"
	"weblogindemo/errno"
	"weblogindemo/middleware"
	"weblogindemo/model"
)

// 注册用户
func RegisterUser(data model.User) errno.RET {
	status := CheckUser(data.Username)
	if status.RetCode() == errno.OK.RetCode() { //用户名不存在,将用户名写入数据库
		// 保存数据库之前需要对密码加密
		//data.Password = ScryptPasswd(data.Password)
		err := Dao.DB.Create(&data).Error
		if err != nil {
			return errno.ERROR
		}
	}
	return status
}

// 用户登陆
func LoginUser(data model.User) (errno.RET, string) {
	var (
		token  string
		status errno.RET
	)
	status = CheckLogin(data.Username, data.Password)
	if status.RetCode() == errno.OK.RetCode() {
		token, status = middleware.SetToken(data.Username)
	}
	return status, token
}

// 查询用户是否存在
func CheckUser(userName string) errno.RET {
	var data model.User
	Dao.DB.Select("id").Where("username = ?", userName).First(&data)
	if data.ID > 0 {
		return errno.ERROR_USERNAME_USED
	}
	return errno.OK
}

// 登录验证
func CheckLogin(userName string, passWord string) errno.RET {
	var user model.User
	Dao.DB.Where("username = ?", userName).First(&user)
	if user.ID == 0 {
		return errno.ERROR_USER_NOT_EXIST
	}
	if encry.ScryptPasswd(passWord) != user.Password {
		return errno.ERROR_PASSWORD_WRONG
	}
	return errno.OK
}
