package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/leebenson/conform"
	"poc/internal/logging"
	"poc/internal/use_case/user"
	"poc/web/api/handlers"
	handlersV1 "poc/web/api/handlers/v1"
	"poc/web/api/middlewares"
	"reflect"
	"strings"
)

type server struct {
	pm       persistenceManager
	echo     *echo.Echo
	enforcer *casbin.Enforcer
}

type customBinder struct{}

type CustomValidator struct {
	validator  *validator.Validate
	translator *ut.UniversalTranslator
}

func newServer(pm persistenceManager) (server, error) {
	enforcer, err := casbinEnforcer()
	if err != nil {
		return server{}, err
	}
	return server{
		pm:       pm,
		echo:     createEchoInstance(),
		enforcer: enforcer,
	}, nil
}

func (s server) start(port string) error {
	s.registerRoutes()
	return s.echo.Start(fmt.Sprintf(":%s", port))
}

func createEchoInstance() *echo.Echo {
	http := echo.New()
	http.Use(echoMiddleware.RequestID())
	http.Use(middlewares.ContextLogger)
	http.Use(middlewares.Logger)
	http.Validator = buildCustomValidator()
	http.Binder = echo.Binder(customBinder{})

	prom := prometheus.NewPrometheus("echo", nil)
	prom.Use(http)

	return http
}

func (cb customBinder) Bind(i interface{}, c echo.Context) (err error) {
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != nil {
		return err
	}

	return conform.Strings(i)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		return logging.NewValidationError(err, cv.translator)
	}

	return nil
}

func buildCustomValidator() *CustomValidator {
	v := validator.New()
	if err := v.RegisterValidation("notblank", validators.NotBlank); err != nil {
		return nil
	}
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	defaultLang := en.New()
	uniTranslator := ut.New(defaultLang, defaultLang)

	defaultTrans, _ := uniTranslator.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v, defaultTrans)

	return &CustomValidator{
		validator:  v,
		translator: uniTranslator,
	}
}

func casbinEnforcer() (*casbin.Enforcer, error) {
	enforcer, err := casbin.NewEnforcer("./resources/auth.conf", "./resources/policy.csv")
	if err != nil {
		return nil, err
	}

	return enforcer, nil
}

func (s server) registerRoutes() {
	authMiddleware := middlewares.NewAuthMiddleware(s.pm.userRepository, s.enforcer)
	api := s.echo.Group("/api")
	{
		api.GET("/health", handlers.Health())
		api.GET("/metrics", handlers.Metrics())
	}

	v1 := api.Group("/v1")
	v1.Use(authMiddleware.Auth)
	{
		userHandler := v1.Group("/users")
		{
			userHandler.GET("", handlersV1.ListUsers(user.NewFindAllUsers(s.pm.userRepository)))
			userHandler.POST("", handlersV1.CreateUser(user.NewCreateUser(s.pm.userRepository)))
			userHandler.GET("/:id", handlersV1.FindUserById(user.NewFindUserById(s.pm.userRepository)))
			userHandler.PUT("/:id", handlersV1.UpdateUser(user.NewUpdateUser(s.pm.userRepository)))
			userHandler.DELETE("/:id", handlersV1.DeleteUSer(user.NewDeleteUser(s.pm.userRepository)))
		}
	}
}
