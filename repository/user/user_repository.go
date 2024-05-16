package user_repository

import (
	"context"

	user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"
)

type UserRepository interface {
	Register(ctx context.Context, req user_entity.User) (user_entity.UserData, error)
	Edit(ctx context.Context, req user_entity.User) error
	Login(ctx context.Context, req user_entity.User) (user_entity.User, error)
	GiveAccess(ctx context.Context, req user_entity.User) (user_entity.User, error)
	Search(ctx context.Context, req user_entity.UserGetRequest) (*[]user_entity.UserResponseData, error)
	Delete(ctx context.Context, userId string) (user_entity.User, error)
}
