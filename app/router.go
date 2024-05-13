package app

import (
	"fmt"

	"github.com/kangman53/project-sprint-halo-suster/controller"
	"github.com/kangman53/project-sprint-halo-suster/helpers"

	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	user_repository "github.com/kangman53/project-sprint-halo-suster/repository/user"
	auth_service "github.com/kangman53/project-sprint-halo-suster/service/auth"
	user_service "github.com/kangman53/project-sprint-halo-suster/service/user"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterBluePrint(app *fiber.App, dbPool *pgxpool.Pool) {
	validator := validator.New()
	// register custom validator
	helpers.RegisterCustomValidator(validator)

	authService := auth_service.NewAuthService()

	userRepository := user_repository.NewUserRepository(dbPool)
	userService := user_service.NewUserService(userRepository, authService, validator)
	userController := controller.NewUserController(userService)

	// Users API
	userApi := app.Group("/v1/user")
	userApi.Post("/it/register", userController.Register)
	userApi.Post("/it/login", userController.Login)
	userApi.Post("/nurse/login", userController.Login)

	// JWT middleware
	// app.Use(helpers.CheckTokenHeader)
	app.Use(helpers.GetTokenHandler())

	// Nurse Management Middleware that requires "it" access
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Locals("userRole"))
		if userRole := c.Locals("userRole"); userRole != "it" {
			return exc.ForbiddenException("Access Forbidden")
		}
		return c.Next()

	})
	userApi.Get("/", userController.Get)
	userApi.Post("/nurse/register", userController.Register)
	userApi.Post("/nurse/:userId/access", userController.GiveAccess)
}
