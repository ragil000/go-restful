package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ragil000/go-restful.git/helpers"
)

// AuthenticationAPIKey validates the API KEY, return 403 if not valid
func AuthenticationAPIKey() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.GetHeader("X-API-KEY") == "" {
			response := helpers.BuildErrorResponse("Failed to process request", "API Key not found", nil)
			context.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		} else if context.GetHeader("X-API-KEY") != os.Getenv("X_API_KEY") {
			response := helpers.BuildErrorResponse("Failed to process request", "API Key invalid", nil)
			context.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}
	}
}
