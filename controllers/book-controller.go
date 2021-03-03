package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ragil000/go-restful.git/dto"
	"github.com/ragil000/go-restful.git/entities"
	"github.com/ragil000/go-restful.git/helpers"
	"github.com/ragil000/go-restful.git/services"
)

// BookController is a...
type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService services.BookService
	jwtService  services.JWTService
}

// NewBookController is a...
func NewBookController(bookServ services.BookService, jwtServ services.JWTService) BookController {
	return &bookController{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
}

func (c *bookController) All(context *gin.Context) {
	var books []entities.Book = c.bookService.All()
	res := helpers.BuildResponse(true, "OK", books)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	var book entities.Book = c.bookService.FindByID(id)
	if (book == entities.Book{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObject{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", book)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Insert(context *gin.Context) {
	fmt.Println(context.GetPostForm("title"))
	var bookCreatedDTO dto.BookCreateDTO
	fmt.Println(bookCreatedDTO)
	errDTO := context.ShouldBind(&bookCreatedDTO)
	fmt.Println(errDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bookCreatedDTO.UserID = convertedUserID
		}
		result := c.bookService.Insert(bookCreatedDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *bookController) Update(context *gin.Context) {
	fmt.Println("Masuk kok")
	var bookUpdateDTO dto.BookUpdateDTO
	fmt.Println("Masuk kok 0")
	errDTO := context.ShouldBind(&bookUpdateDTO)
	fmt.Println(errDTO)
	fmt.Println("Masuk kok 1")
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObject{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println("Masuk kok 2")
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	fmt.Print("userID: ")
	fmt.Print(userID)
	bookID, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to get id", "No param id were found", helpers.EmptyObject{})
		context.JSON(http.StatusBadRequest, response)
	}
	bookUpdateDTO.ID = bookID
	if c.bookService.IsAllowedToEdit(userID, bookID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObject{})
		context.JSON(http.StatusBadRequest, response)
	}
}

func (c *bookController) Delete(context *gin.Context) {
	var book entities.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to get id", "No param id were found", helpers.EmptyObject{})
		context.JSON(http.StatusBadRequest, response)
	}
	book.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObject{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObject{})
		context.JSON(http.StatusBadRequest, response)
	}
}

func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
