package mysql

import (
	"github.com/jmoiron/sqlx"
	"pocket-book/comm"
)

type SqlUtil struct {
}

func (SqlUtil) NewSession() *sqlx.Tx {
	session, err := db.Beginx()
	if err != nil {
		comm.MysqlLogger.Error().Msgf("new session failed: %s", err.Error())
		return nil
	}
	return session
}
