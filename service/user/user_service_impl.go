package user_service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	user_entity "github.com/kangman53/project-sprint-halo-suster/entity/user"
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	helpers "github.com/kangman53/project-sprint-halo-suster/helpers"
	userRep "github.com/kangman53/project-sprint-halo-suster/repository/user"
	authService "github.com/kangman53/project-sprint-halo-suster/service/auth"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	UserRepository userRep.UserRepository
	AuthService    authService.AuthService
	Validator      *validator.Validate
}

func NewUserService(userRepository userRep.UserRepository, authService authService.AuthService, validator *validator.Validate) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
		AuthService:    authService,
		Validator:      validator,
	}
}

func (service *userServiceImpl) Register(ctx *fiber.Ctx, req user_entity.UserRegisterRequest) (user_entity.UserResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(req); err != nil {
		return user_entity.UserResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	role := strings.Split(ctx.OriginalURL(), "/")[3]
	user := user_entity.User{
		Name: req.Name,
		Nip:  strconv.Itoa(req.Nip),
		Role: role,
	}
	var nipType string
	if role == "it" {
		if err := service.Validator.Var(req.Password, "min=5,max=33"); err != nil {
			return user_entity.UserResponse{}, exc.BadRequestException("Bad request: Invalid Password")
		}
		hashPassword, err := helpers.HashPassword(req.Password)
		if err != nil {
			return user_entity.UserResponse{}, err
		}
		user.Password = hashPassword
		nipType = "nipIT"
	} else {
		if err := service.Validator.Var(req.IdentityCardScanImg, "validateUrl"); err != nil {
			return user_entity.UserResponse{}, exc.BadRequestException("Bad request: Invalid IdentityCardScanImg")
		}
		nipType = "nipNurse"
		user.IdentityCardScanImg = req.IdentityCardScanImg
	}

	if err := service.Validator.Var(user.Nip, nipType); err != nil {
		return user_entity.UserResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: Invalid NIP for %s role", role))
	}

	userContext := ctx.UserContext()
	userRegistered, err := service.UserRepository.Register(userContext, user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return user_entity.UserResponse{}, exc.ConflictException("User with this nip already registered")
		}
		return user_entity.UserResponse{}, err
	}

	token, err := service.AuthService.GenerateToken(userContext, userRegistered.Id, req.Role)
	if err != nil {
		return user_entity.UserResponse{}, err
	}

	userRegistered.AccessToken = token
	userRegistered.Name = req.Name
	userRegistered.Nip = req.Nip
	return user_entity.UserResponse{
		Message: "User registered",
		Data:    &userRegistered,
	}, nil
}

func (service *userServiceImpl) Login(ctx *fiber.Ctx, req user_entity.UserLoginRequest) (user_entity.UserResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(req); err != nil {
		return user_entity.UserResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	nipType := "nipIT"
	user := user_entity.User{
		Nip: strconv.Itoa(req.Nip),
	}
	role := strings.Split(ctx.OriginalURL(), "/")[3]
	if role != "it" {
		nipType = "nipNurse"
	}

	if err := service.Validator.Var(user.Nip, nipType); err != nil {
		return user_entity.UserResponse{}, exc.NotFoundException("User is not found")
	}

	userContext := ctx.UserContext()
	userLogin, err := service.UserRepository.Login(userContext, user)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return user_entity.UserResponse{}, exc.NotFoundException("User is not found")
		}

		return user_entity.UserResponse{}, err
	}

	if userLogin.Password == "" {
		return user_entity.UserResponse{}, exc.BadRequestException("User is not having access")
	}

	if _, err = helpers.ComparePassword(userLogin.Password, req.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return user_entity.UserResponse{}, exc.BadRequestException("Invalid password")
		}

		return user_entity.UserResponse{}, err
	}

	token, err := service.AuthService.GenerateToken(userContext, userLogin.Id, userLogin.Role)
	if err != nil {
		return user_entity.UserResponse{}, err
	}

	return user_entity.UserResponse{
		Message: "User logged successfully",
		Data: &user_entity.UserData{
			Id:          userLogin.Id,
			Nip:         req.Nip,
			Name:        userLogin.Name,
			AccessToken: token,
		},
	}, nil
}

func (service *userServiceImpl) GiveAccess(ctx *fiber.Ctx, req user_entity.NurseAccessRequest) (user_entity.UserResponse, error) {
	// validate by rule we defined in _request_entity.go
	if err := service.Validator.Struct(req); err != nil {
		return user_entity.UserResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	userId := ctx.Params("userId")
	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return user_entity.UserResponse{}, err
	}
	user := user_entity.User{
		Id:       userId,
		Password: hashPassword,
		Role:     "nurse",
	}

	userContext := ctx.UserContext()
	userLogin, err := service.UserRepository.GiveAccess(userContext, user)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return user_entity.UserResponse{}, exc.NotFoundException("User is not found")
		}

		return user_entity.UserResponse{}, err
	}

	nip, _ := strconv.Atoi(userLogin.Nip)
	return user_entity.UserResponse{
		Message: "Successfully create access for nurse",
		Data: &user_entity.UserData{
			Id:   userLogin.Id,
			Name: userLogin.Name,
			Nip:  nip,
		},
	}, nil

}

func (service *userServiceImpl) Search(ctx *fiber.Ctx, req user_entity.UserGetRequest) (user_entity.UserGetResponse, error) {
	if err := service.Validator.Struct(req); err != nil {
		return user_entity.UserGetResponse{}, exc.BadRequestException(fmt.Sprintf("Bad request: %s", err))
	}

	userSearch, err := service.UserRepository.Search(ctx.UserContext(), req)
	if err != nil {
		return user_entity.UserGetResponse{}, err
	}

	return user_entity.UserGetResponse{
		Message: "Successfully retrieved users",
		Data:    userSearch,
	}, nil

}

func (service *userServiceImpl) Delete(ctx *fiber.Ctx) (user_entity.UserResponse, error) {
	userId := ctx.Params("userId")
	deletedUser, err := service.UserRepository.Delete(ctx.Context(), userId)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return user_entity.UserResponse{}, exc.NotFoundException("User is not found")
		}
		return user_entity.UserResponse{}, err
	}
	nip, _ := strconv.Atoi(deletedUser.Nip)
	return user_entity.UserResponse{
		Message: "Successfully deleted user",
		Data: &user_entity.UserData{
			Id:   deletedUser.Id,
			Name: deletedUser.Name,
			Nip:  nip,
		},
	}, nil
}
