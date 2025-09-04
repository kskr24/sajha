package main

import (
	"net/http"
	"time"

	"github.com/kskr24/sajha/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type App struct {
	Q *db.Queries
	L *zap.Logger
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ZapLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			stop := time.Now()
			latency := stop.Sub(start)

			logger.Info("request",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
				zap.Int("status", c.Response().Status),
				zap.String("ip", c.RealIP()),
				zap.Duration("latency", latency))
			return err
		}
	}
}

func (app *App) loginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var u User
		if err := c.Bind(&u); err != nil {
			return c.JSON(http.StatusBadRequest,
				map[string]string{"error": "invalid request"})
		}
		app.Q.InsertUser(c.Request().Context(), db.InsertUserParams{
			Name:     u.Username,
			Password: u.Password,
		})
		return c.JSON(http.StatusOK,
			map[string]string{"token": "tok123"})
	}
}
func (app *App) listUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	}
}

func (app *App) helloHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {

			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello! Authenticated User."})
	}
}

func main() {
	pool := db.Init()
	q := db.New(pool)

	defer pool.Close()
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:4200"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}},
	))
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	app := &App{
		Q: q, L: logger,
	}
	e.Use(ZapLogger(logger))

	e.POST("/v1/login", app.loginHandler())
	e.GET("/v1/users", app.listUserHandler())
	e.GET("/v1/hello", app.helloHandler())

	e.Logger.Fatal(e.Start(":8080"))
}
