package mysql

import (
	sql2 "database/sql"
	"errors"
	"fmt"
	"pocket-book/comm"
	"pocket-book/models"
	"strings"
)

func AddExpenses(reqData *models.ParmaAddExpenses, userId int) (err error) {
	session := SqlUtil{}.NewSession()
	sql := `insert into t_transactions
				(user_id, category_id, title, description, amount, transaction_date, type)
			values (?, ?, ?, ?, ?, ?, ?);`
	_, err = session.Exec(sql, userId, reqData.CategoryId, reqData.Title, reqData.Description, reqData.Amount, reqData.TransactionDate, reqData.Type)
	err = SqlUtil{}.ExecOpt(err, session)
	return err
}

func EditExpenses(reqData *models.ParmaEditExpenses, userId int) (err error) {
	session := SqlUtil{}.NewSession()
	sql := `
		update t_transactions
		set title            = ?,
			description      = ?,
			amount= ?,
			category_id      = ?,
			transaction_date = ?,
			type             = ?
		where id = ?;
		`
	_, err = session.Exec(sql, reqData.Title, reqData.Description, reqData.Amount, reqData.CategoryId, reqData.TransactionDate, reqData.Type, reqData.BillId)
	err = SqlUtil{}.ExecOpt(err, session)
	comm.MysqlLogger.Info().Msgf("%s 编辑了记账信息！", userId)
	return err
}

func DeleteExpenses(billId, userId int) (err error) {
	session := SqlUtil{}.NewSession()
	sql := `delete from t_transactions where id = ?;`
	_, err = session.Exec(sql, billId)
	err = SqlUtil{}.ExecOpt(err, session)
	comm.MysqlLogger.Info().Msgf("%s 删除了记账信息！", userId)
	return err
}

func SearchCommExpenses(reqData *models.ParamSearchExpenses, userId int) (err error, results *models.ResponseSearchPay) {
	results = new(models.ResponseSearchPay)
	sql := `
			SELECT tt.id as bill_id,
			       DATE_FORMAT(tt.transaction_date, '%Y-%m-%d %H:%i:%s') as date,
				   title,
				   tt.description,
				   amount,
				   tc.name                                               as cate
			FROM t_transactions tt
					 JOIN t_categories tc ON tc.id = tt.category_id
			WHERE tt.user_id = ?
			  AND tt.type = ?
			`
	cntSql := `
			SELECT count(tt.id) as total                                            
			FROM t_transactions tt
					 JOIN t_categories tc ON tc.id = tt.category_id
			WHERE tt.user_id = ?
			  AND tt.type = ?
			`
	// 初始化参数切片
	args := []interface{}{userId, reqData.Type}
	// 构建额外的 WHERE 条件
	var whereClauses []string
	var values []interface{}

	if reqData.Title != "" {
		whereClauses = append(whereClauses, "tt.title LIKE concat('%', ?, '%')")
		values = append(values, reqData.Title)
	}
	if reqData.TransactionBeginDate != "" && reqData.TransactionEndDate != "" {
		whereClauses = append(whereClauses, "tt.transaction_date BETWEEN ? AND ?")
		values = append(values, reqData.TransactionBeginDate, reqData.TransactionEndDate)
	}

	// 如果有额外的 WHERE 条件，添加到 SQL 语句中
	if len(whereClauses) > 0 {
		filterSql := " AND " + strings.Join(whereClauses, " AND ")
		sql += filterSql
		cntSql += filterSql
		args = append(args, values...) // 合并参数
	}
	pageSql := fmt.Sprintf("limit %d offset %d ;", reqData.Limit, reqData.Offset)
	sql += " order by tt.transaction_date desc " + pageSql

	// 执行查询
	err = db.Select(&results.ResList, sql, args...)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			comm.MysqlLogger.Info().Msg("No rows found")
			return
		}
		comm.MysqlLogger.Error().Msg(err.Error())
		err = comm.ErrServerBusy
		return
	}
	_ = db.Get(&results.Total, cntSql, args...)

	return
}

func SearchExpensesPreview(reqData *models.ParamSearchExpensesPreview, userId int) (result *models.ResponseSearchTotalPayIncome, err error) {
	result = new(models.ResponseSearchTotalPayIncome)
	sql := `
		select COALESCE(sum(if(tt.type = 1, tt.amount, 0)), 0)                                                   as total_pay,
			   COALESCE(sum(if(tt.type = 2, tt.amount, 0)), 0)                                                   as total_income,
			   COALESCE(sum(if(tt.type = 2, tt.amount, 0)), 0) - COALESCE(sum(if(tt.type = 1, tt.amount, 0)), 0) as overall
		from t_transactions tt
		where tt.user_id = ? 
	`
	args := []interface{}{userId}

	if reqData.StartTime != "" {
		sql += " and tt.transaction_date BETWEEN ? AND ? ;"
		args = append(args, reqData.StartTime, reqData.EndTime)
	}
	err = db.Get(result, sql, args...)
	if err != nil {
		return nil, err
	}

	return
}
