package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ragil000/go-restful.git/dto"
	"github.com/ragil000/go-restful.git/entities"
	"github.com/ragil000/go-restful.git/helpers"
	"github.com/ragil000/go-restful.git/services"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(context *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := context.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entities.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10))
		v.Token = generateToken
		response := helpers.BuildResponse(true, "OK!", v)
		context.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse("Please check again your credential", "Invalid credential", helpers.EmptyObject{})
	context.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(context *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := context.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObject{})
		context.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		createdUser.Token = token
		response := helpers.BuildResponse(true, "OK!", createdUser)
		context.JSON(http.StatusCreated, response)
	}
}
