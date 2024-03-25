package middleware

import (
	"net/http"
	"DENTAL-CLINIC/pkg/web"
	"github.com/gin-gonic/gin"
	"os"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusBadRequest, web.NewBadRequestApiError("No se ha encontrado el token"))
			c.Abort()
			return
		}

		tokenFromEnv := os.Getenv("TOKEN")
		if token != tokenFromEnv {
			c.JSON(http.StatusBadRequest, web.NewBadRequestApiError("Token invalido"))
			c.Abort()
			return
		}

		c.Next()
	}
}