package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ragil000/go-restful.git/config"
	"github.com/ragil000/go-restful.git/controllers"
	"github.com/ragil000/go-restful.git/middleware"
	"github.com/ragil000/go-restful.git/repositories"
	"github.com/ragil000/go-restful.git/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                    = config.Connection()
	userRepository repositories.UserRepository = repositories.NewUserRespository(db)
	bookRepository repositories.BookRepository = repositories.NewBookRepository(db)
	jwtService     services.JWTService         = services.NewJWTService()
	authService    services.AuthService        = services.NewAuthService(userRepository)
	userService    services.UserService        = services.NewUserService(userRepository)
	bookService    services.BookService        = services.NewBookService(bookRepository)
	authController controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController  = controllers.NewUserController(userService, jwtService)
	bookController controllers.BookController  = controllers.NewBookController(bookService, jwtService)
)

func main() {
	// defer config.CloseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth", middleware.AuthenticationAPIKey())
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	bookRoutes := r.Group("api/book", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}

	r.Run()
}
