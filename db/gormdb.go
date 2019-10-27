package db

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type GormDB struct {
	*gorm.DB
	gdbDone bool
}

func (c *GormDB) DbCommit() error {
	if c.gdbDone {
		return nil
	}
	tx := c.Commit()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		return err
	}
	return nil
}

func (c *GormDB) DbRollback() error {
	if c.gdbDone {
		return nil
	}
	tx := c.Rollback()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		return err
	}
	return nil
}
