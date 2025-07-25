package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		FindOneByUsername(ctx context.Context, username string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (user *customUserModel) FindOneByUsername(ctx context.Context, username string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, user.table)

	var resp User
	err := user.conn.QueryRowCtx(ctx, &resp, query, username)

	switch err {
	case nil:
		return &resp, nil
	case sql.ErrNoRows, sqlx.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}
