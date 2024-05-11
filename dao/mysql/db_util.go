package mysql

import (
	"database/sql"
	"pocket-book/comm"
)

type SqlUtil struct {
}

func (SqlUtil) NewSession() *sql.Tx {
	session, err := db.Begin()
	if err != nil {
		comm.MysqlLogger.Error().Msgf("new session failed: %s", err.Error())
		return nil
	}
	return session
}
