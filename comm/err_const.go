package comm

import "errors"

var (
	ErrUserNotExist    = errors.New("用户不存在")
	ErrUserExist       = errors.New("用户已经存在")
	ErrPswUName        = errors.New("用户名或密码输入错误")
	ErrInvalidPswUName = errors.New("用户名或密码不合法")
	ErrServerBusy      = errors.New("服务器繁忙")
)
