package repository_import

import (
	"salei/app/internal/repository"
	"salei/app/internal/repository/postgres"
)

type PostgresRepo struct {
	repository.Auth
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{
		postgres.NewAuthRepo(),
	}
}
