package mysql

import (
	"pocket-book/models"
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
