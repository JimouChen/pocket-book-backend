package mysql

import (
	sql2 "database/sql"
	"errors"
	"pocket-book/comm"
	"pocket-book/models"
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

func EditCategoryById(id int, name string) (err error) {
	session := SqlUtil{}.NewSession()
	if session == nil {
		return comm.ErrCreateMysqlSession
	}
	sql := "update t_categories set name = ? where id = ?;"
	_, err = session.Exec(sql, name, id)
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

func SearchCategoryByUsername(username string) (err error, results []*models.ParamEditCategory) {
	sql := "select tc.id, name from t_user2cate tu2c join t_categories tc on tu2c.category_id = tc.id  join t_users on t_users.id = tu2c.user_id where t_users.username = ?;"

	// 执行查询
	if err = db.Select(&results, sql, username); err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			comm.MysqlLogger.Info().Msg("该用户无设置新增分类")
			return
		}
		comm.MysqlLogger.Error().Msg(err.Error())
		err = comm.ErrServerBusy
		return
	}
	return
}

func SearchAllCategory() (err error, results []*models.ParamEditCategory) {
	sql := "select id, name from t_categories;"

	if err = db.Select(&results, sql); err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			comm.MysqlLogger.Info().Msg("系统还未设置分类，请先添加分类！")
			return
		}
		comm.MysqlLogger.Error().Msg(err.Error())
		err = comm.ErrServerBusy
		return
	}
	return
}
