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

func (SqlUtil) ExecQueries(tx *sqlx.Tx, queries []string, args [][]interface{}) error {
	for i, query := range queries {
		reboundQuery := tx.Rebind(query)
		if _, err := tx.Exec(reboundQuery, args[i]...); err != nil {
			return err
		}
	}
	return nil
}

func (SqlUtil) ExecOpt(err error, session *sqlx.Tx) error {
	if err != nil {
		_ = session.Rollback()
		return comm.ErrServerBusy
	}

	if err = session.Commit(); err != nil {
		_ = session.Rollback()
		comm.MysqlLogger.Error().Msgf("rollback session:%s", err.Error())
		return comm.ErrServerBusy
	}
	return nil
}
