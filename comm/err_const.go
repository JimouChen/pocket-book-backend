package comm

import "errors"

var (
	ErrUserNotExist       = errors.New("用户不存在")
	ErrUserExist          = errors.New("用户已经存在")
	ErrCategoryExist      = errors.New("分类已经存在")
	ErrUseCateNotExist    = errors.New("用户自己分类不存在")
	ErrPswUName           = errors.New("用户名或密码输入错误")
	ErrInvalidPswUName    = errors.New("用户名或密码不合法")
	ErrServerBusy         = errors.New("服务器繁忙")
	ErrReadMysql          = errors.New("mysql数据库读取错误")
	ErrCreateMysqlSession = errors.New("mysql创建session失败")
)
