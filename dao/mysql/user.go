package mysql

import (
	"pocket-book/comm"
	"pocket-book/models"
)

func CheckUserIsExist(username string) (err error) {
	sql := "select count(id) from t_users where username = ?;"
	var cnt int
	if err := db.Get(&cnt, sql, username); err != nil {
		return comm.ErrServerBusy
	}
	if cnt > 0 {
		return comm.ErrUserExist
	}
	return
}

func InsertUser(user *models.ParamUser) (err error) {
	session, err := db.Begin()
	if err != nil {
		comm.MysqlLogger.Error().Msgf("InsertUser failed: %s", err.Error())
	}
	//加密密码
	//user.Password = Md5Psw(user.Password)
	//写入数据库
	sql := "insert t_users(username, password) value( ?, ?);"

	_, err = session.Exec(sql, user.Username, user.Password)
	if err != nil {
		_ = session.Rollback()
		return comm.ErrServerBusy
	}
	if err = session.Commit(); err != nil {
		return err
	}
	return
}

func CheckLogin(user *models.ParamUser) (UserId int, err error) {
	sql := "select id from t_users where username = ? and password = ? limit 1;"
	var resId int
	if err = db.Get(&resId, sql, user.Username, user.Password); err != nil {
		comm.MysqlLogger.Error().Msg(err.Error())
		return 0, comm.ErrReadMysql
	}
	// 记录最近的登陆时间
	UpdateLoginTimeSql := "update t_users set last_login_time = now() where id = ?;"
	session, err := db.Begin()
	if _, err = session.Exec(UpdateLoginTimeSql, resId); err != nil {
		_ = session.Rollback()
		return 0, comm.ErrServerBusy
	}
	if err = session.Commit(); err != nil {
		return 0, err
	}
	comm.MysqlLogger.Debug().Msgf("用户：%s 登陆成功！", user.Username)
	return resId, nil
}
