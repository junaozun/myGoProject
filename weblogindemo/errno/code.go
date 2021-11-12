package errno

var (
	// OK
	OK    = NewError(0, "OK")
	ERROR = NewError(1, "ERROR")

	// 模块级错误码 - 用户模块
	ERROR_USERNAME_USED   = NewError(10001, "用户名已存在!")
	ERROR_PASSWORD_WRONG  = NewError(10002, "密码错误")
	ERROR_USER_NOT_EXIST  = NewError(10003, "用户不存在")
	ERROR_TOKEN_Not_EXIST = NewError(10004, "TOKEN不存在")
	ERROR_TOKEN_RUNTIME   = NewError(10005, "TOKEN已过期")
	ERROR_TOKEN_WRONG     = NewError(10006, "TOKEN不正确")
	ERROR_TOKEN_TYPE_ERR  = NewError(10007, "TOKEN格式错误")

	// ...
)
