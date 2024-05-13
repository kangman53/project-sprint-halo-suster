package user_repository

import (
	"context"

	user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepositoryImpl struct {
	DBpool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) UserRepository {
	return &userRepositoryImpl{
		DBpool: dbPool,
	}
}

func (repository *userRepositoryImpl) Register(ctx context.Context, user user_entity.User) (user_entity.UserData, error) {
	var userId string
	query := "INSERT INTO users (name, nip, role, password) VALUES ($1, $2, $3, $4) RETURNING id"
	if err := repository.DBpool.QueryRow(ctx, query, user.Name, user.Nip, user.Role, user.Password).Scan(&userId); err != nil {
		return user_entity.UserData{}, err
	}

	return user_entity.UserData{Id: userId}, nil
}

func (repository *userRepositoryImpl) Login(ctx context.Context, user user_entity.User) (user_entity.User, error) {
	query := "SELECT id, name, password, role FROM users WHERE nip = $1 LIMIT 1"
	row := repository.DBpool.QueryRow(ctx, query, user.Nip)

	var loggedInUser user_entity.User
	err := row.Scan(&loggedInUser.Id, &loggedInUser.Name, &loggedInUser.Password, &loggedInUser.Role)
	if err != nil {
		return user_entity.User{}, err
	}

	return loggedInUser, nil
}
