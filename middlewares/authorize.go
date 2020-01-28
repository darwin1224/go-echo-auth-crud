package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Authorize(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			credentials := c.Get("user").(*jwt.Token)
			claims := credentials.Claims.(jwt.MapClaims)
			roleUser := claims["role_user"]
			access := false
			for _, role := range roles {
				if strings.ToLower(role) == strings.ToLower(roleUser.(string)) {
					access = true
				}
			}
			if !access {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "You are not allowed to access this route"})
			}
			return next(c)
		}
	}
}
