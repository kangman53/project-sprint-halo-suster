package user_repository

import (
	"context"

	user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"
)

type UserRepository interface {
	Register(ctx context.Context, req user_entity.User) (user_entity.UserData, error)
	Login(ctx context.Context, req user_entity.User) (user_entity.User, error)
}
