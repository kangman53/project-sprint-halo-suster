package user_service

import (
	"github.com/gofiber/fiber/v2"
	user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"
)

type UserService interface {
	Register(ctx *fiber.Ctx, req user_entity.UserRegisterRequest) (user_entity.UserResponse, error)
	Login(ctx *fiber.Ctx, req user_entity.UserLoginRequest) (user_entity.UserResponse, error)
	GiveAccess(ctx *fiber.Ctx, req user_entity.NurseAccessRequest) (user_entity.UserResponse, error)
	Search(ctx *fiber.Ctx, req user_entity.UserGetRequest) (user_entity.UserGetResponse, error)
	Delete(ctx *fiber.Ctx) (user_entity.UserResponse, error)
}
