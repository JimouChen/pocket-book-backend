package mysql

import (
	sql2 "database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"pocket-book/comm"
	"pocket-book/models"
)

func CheckCategoryIsExist(name string, userId int) (error, int, int) {
	sql := "select count(id) from t_categories where name = ?;"
	subSql := `
		select count(tc.id) as cnt
		from t_categories tc
				 join t_user2cate tu2c on tu2c.category_id = tc.id
		where name = ? and tu2c.user_id = ?;
		`
	var cnt, subCnt int
	if err := db.Get(&cnt, sql, name); err != nil {
		return comm.ErrServerBusy, 0, 0
	}
	if err := db.Get(&subCnt, subSql, name, userId); err != nil {
		return comm.ErrServerBusy, 0, 0
	}

	return nil, cnt, subCnt
}

func AddCategory(name string, userId int) (err error) {
	_, cnt, subCnt := CheckCategoryIsExist(name, userId)
	if (cnt > 0) && (subCnt > 0) {
		return comm.ErrCategoryExist
	}
	session := SqlUtil{}.NewSession()
	if session == nil {
		return comm.ErrCreateMysqlSession
	}

	if (cnt > 0) && (subCnt == 0) {
		getCateIdSql := "select id from t_categories where name = ?;"
		var CateId int
		if err := db.Get(&CateId, getCateIdSql, name); err != nil {
			return comm.ErrServerBusy
		}
		subSql := "insert into t_user2cate (user_id, category_id) VALUES (?, ?);"
		_, err = session.Exec(subSql, userId, CateId)
		if err != nil {
			_ = session.Rollback()
			return comm.ErrServerBusy
		}
		_ = session.Commit()
		return
	}
	// 	cnt == 0 && subCnt == 0
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

func DeleteCategoryByNames(categoryNames []string, userId int) (err error) {
	session := SqlUtil{}.NewSession()

	queries := []string{
		"DELETE FROM t_transactions WHERE category_id IN (SELECT id FROM t_categories tc WHERE tc.name IN (?)) AND user_id = ?",
		"DELETE FROM t_user2cate WHERE category_id IN (SELECT id FROM t_categories tc WHERE tc.name IN (?)) AND user_id = ?",
	}

	args := [][]interface{}{
		{categoryNames, userId},
		{categoryNames, userId},
	}
	for i, query := range queries {
		queries[i], args[i], err = sqlx.In(query, args[i]...)
		if err != nil {
			_ = session.Rollback()
			return err
		}
	}
	sqlUtil := SqlUtil{}
	if err := sqlUtil.ExecQueries(session, queries, args); err != nil {
		_ = session.Rollback()
		return err
	}

	if err := session.Commit(); err != nil {
		_ = session.Rollback()
		return err
	}

	return nil
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
