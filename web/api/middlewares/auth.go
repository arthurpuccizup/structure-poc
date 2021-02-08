package middlewares

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"poc/internal/repository"
	"poc/internal/tracking"
	"strings"
)

type AuthMiddleware struct {
	userRepo repository.UserRepository
	enforcer *casbin.Enforcer
}

type AuthToken struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewAuthMiddleware(userRepository repository.UserRepository, enforcer *casbin.Enforcer) AuthMiddleware {
	return AuthMiddleware{
		userRepo: userRepository,
		enforcer: enforcer,
	}
}

//Could use the user that came from the JWT to Query the Database for the permissions and after
//do a check in a more complex way
func (a AuthMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		openedToken, err := extractToken(c.Request().Header.Get("Authorization"))
		if err != nil {
			return c.JSON(http.StatusForbidden, tracking.NewError("Token not informed or invalid", err, nil))
		}
		fmt.Print(openedToken)

		allowed, enforcerErr := a.enforcer.Enforce("write", c.Request().URL.Path, c.Request().Method)
		if enforcerErr != nil {
			return c.JSON(http.StatusForbidden, tracking.NewError("Token not informed or invalid", enforcerErr, nil))
		}

		if !allowed {
			return c.JSON(http.StatusUnauthorized, tracking.NewError("Not allowed", nil, nil))
		}

		return next(c)
	}
}

func extractToken(authorization string) (AuthToken, error) {
	rToken := strings.TrimSpace(authorization)
	if rToken == "" {
		return AuthToken{}, tracking.NewError("Extract token error", errors.New("token is require"), nil)
	}

	splitToken := strings.Split(rToken, "Bearer ")

	token, _, err := new(jwt.Parser).ParseUnverified(splitToken[1], &AuthToken{})
	if err != nil {
		return AuthToken{}, tracking.NewError("Extract token error", err, nil)
	}

	return *token.Claims.(*AuthToken), nil
}
