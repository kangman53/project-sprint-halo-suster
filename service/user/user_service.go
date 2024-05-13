package user_service

import (
	"github.com/gofiber/fiber/v2"
	user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"
)

type UserService interface {
	Register(ctx *fiber.Ctx, req user_entity.UserRegisterRequest) (user_entity.UserResponse, error)
	Login(ctx *fiber.Ctx, req user_entity.UserLoginRequest) (user_entity.UserResponse, error)
}
