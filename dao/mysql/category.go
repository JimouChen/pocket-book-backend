package mysql

import (
	"pocket-book/comm"
)

func CheckCategoryIsExist(name string) (err error) {
	sql := "select count(id) from t_categories where name = ?;"
	var cnt int
	if err := db.Get(&cnt, sql, name); err != nil {
		return comm.ErrServerBusy
	}
	if cnt > 0 {
		return comm.ErrCategoryExist
	}
	return
}

func AddCategory(name string, userId int) (err error) {
	session := SqlUtil{}.NewSession()
	if session == nil {
		return comm.ErrCreateMysqlSession
	}
	sql := "insert into t_categories (name) values (?);"
	subSql := "insert into t_user2cate (user_id, category_id) VALUES (?, ?);"
	res, err := session.Exec(sql, name)
	if err != nil {
		_ = session.Rollback()
		return comm.ErrServerBusy
	}
	lastInsertId, err := res.LastInsertId()
	_, err = session.Exec(subSql, userId, lastInsertId)
	if err != nil {
		_ = session.Rollback()
		return comm.ErrServerBusy
	}

	if err = session.Commit(); err != nil {
		_ = session.Rollback()
		comm.MysqlLogger.Error().Msgf("rollback session:%s", err.Error())
		return comm.ErrServerBusy
	}
	return
}

func DeleteCategoryById(cateId int, userId int) (err error) {
	session := SqlUtil{}.NewSession()
	if session == nil {
		return comm.ErrCreateMysqlSession
	}
	subSql := "delete from t_transactions where category_id = ? and user_id = ?;"
	sql := "delete from t_user2cate where category_id = ? and user_id = ?;"
	_, err = session.Exec(subSql, cateId, userId)
	if err != nil {
		_ = session.Rollback()
		comm.MysqlLogger.Error().Msg(err.Error())
		return comm.ErrServerBusy
	}
	_, err = session.Exec(sql, cateId, userId)
	if err != nil {
		_ = session.Rollback()
		comm.MysqlLogger.Error().Msg(err.Error())
		return comm.ErrServerBusy
	}
	if err = session.Commit(); err != nil {
		return
	}
	return
}
