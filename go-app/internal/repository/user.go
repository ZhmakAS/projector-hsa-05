package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"projector-test-app/internal/models"
)

type User interface {
	Save(ctx context.Context, user *models.User) (*models.User, error)
}

type user struct {
	db *sqlx.DB
}

func NewUser(conn *sqlx.DB) User {
	return &user{
		db: conn,
	}
}

func (u *user) Save(ctx context.Context, user *models.User) (*models.User, error) {
	queryBuilder := sq.Insert("users").
		Columns(
			"first_name",
			"last_name",
			"phone",
			"created_at",
		).
		Values(
			user.FirstName,
			user.LastName,
			user.Phone,
			user.CreatedAt,
		)
	sql, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	res, err := u.db.Exec(sql, args...)
	if err != nil {
		return nil, err
	}

	user.Id, err = res.LastInsertId()
	return user, err
}
