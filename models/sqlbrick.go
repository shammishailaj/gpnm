// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).
// Code generated by sqlbrick. DO NOT EDIT IT.

package models

import (
	"github.com/jmoiron/sqlx"
)

// Type definition for SqlBrick. It contains all bricks depends on the number of
// sql files. It also wraps some sqlx func for for convenience.
type SqlBrick struct {
	Db      *sqlx.DB
	PkgInfo *PkgInfoBrick
	User    *UserBrick
}

// NewSqlBrick create a new SqlBrick to operate all bricks.
func NewSqlBrick(db *sqlx.DB) *SqlBrick {
	return &SqlBrick{
		Db:      db,
		PkgInfo: newPkgInfoBrick(db),
		User:    newUserBrick(db),
	}
}