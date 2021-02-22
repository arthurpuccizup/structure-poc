package middlewares

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"poc/internal/logging"
	"poc/internal/repository"
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
	return func(echoCtx echo.Context) error {
		openedToken, err := extractToken(echoCtx.Request().Header.Get("Authorization"))
		if err != nil {
			return echoCtx.JSON(http.StatusForbidden, logging.NewError("Token not informed or invalid", err, nil))
		}
		if logger, ok := logging.LoggerFromContext(echoCtx.Request().Context()); ok {
			logger.Infow(fmt.Sprintf("%v", openedToken))
		}

		allowed, enforcerErr := a.enforcer.Enforce("write", echoCtx.Request().URL.Path, echoCtx.Request().Method)
		if enforcerErr != nil {
			return echoCtx.JSON(http.StatusForbidden, logging.NewError("Token not informed or invalid", enforcerErr, nil))
		}

		if !allowed {
			return echoCtx.JSON(http.StatusUnauthorized, logging.NewError("Not allowed", nil, nil))
		}

		return next(echoCtx)
	}
}

func extractToken(authorization string) (AuthToken, error) {
	rToken := strings.TrimSpace(authorization)
	if rToken == "" {
		return AuthToken{}, logging.NewError("Extract token error", errors.New("token is require"), nil)
	}

	splitToken := strings.Split(rToken, "Bearer ")

	token, _, err := new(jwt.Parser).ParseUnverified(splitToken[1], &AuthToken{})
	if err != nil {
		return AuthToken{}, logging.NewError("Extract token error", err, nil)
	}

	return *token.Claims.(*AuthToken), nil
}
