package mysql

import "pocket-book/comm"

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

func AddCategory(name string) (err error) {
	session := SqlUtil{}.NewSession()
	if session == nil {
		return comm.ErrCreateMysqlSession
	}
	sql := "insert into t_categories (name) values (?);"
	_, err = session.Exec(sql, name)
	if err != nil {
		_ = session.Rollback()
		return comm.ErrServerBusy
	}
	if err = session.Commit(); err != nil {
		return
	}
	return
}
