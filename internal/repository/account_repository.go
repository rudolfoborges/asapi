package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rudolfoborges/asapi/internal/entity"
)

type AccountRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Create(ctx context.Context, account *entity.Account) error
}

type accountRepositoryPostgres struct {
	db *sqlx.DB
}

func NewAccountRepositoryPostgres(db *sqlx.DB) AccountRepository {
	return &accountRepositoryPostgres{
		db: db,
	}
}

func (r *accountRepositoryPostgres) Create(ctx context.Context, account *entity.Account) error {
	sql := `
        INSERT INTO accounts (id, name, email, password, active, created_at, updated_at)
        VALUES (:id, :name, :email, :password, :active, :created_at, :updated_at)
    `

	_, err := r.db.NamedExecContext(ctx, sql, account)
	if err != nil {
		return err
	}

	return nil
}

func (r *accountRepositoryPostgres) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	sql := "SELECT EXISTS(SELECT 1 FROM accounts WHERE email = $1 and active = $2)"

	var exists bool
	err := r.db.GetContext(ctx, &exists, sql, email, true)
	if err != nil {
		return false, err
	}

	return exists, nil
}
