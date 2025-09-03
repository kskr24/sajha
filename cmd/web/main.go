package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	users []User
	mu    sync.Mutex
)

func loginHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "invalid request"})
	}
	mu.Lock()
	users = append(users, *u)
	mu.Unlock()

	return c.JSON(http.StatusOK,
		map[string]string{"token": "tok123"})
}

func listUserHandler(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()
	return c.JSON(http.StatusOK, users)
}

func helloHandler(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {

		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	fmt.Print("NOOOOOO")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello! Authenticated User."})
}
func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:4200"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}},
	))
	e.POST("/v1/login", loginHandler)
	e.GET("/v1/users", listUserHandler)
	e.GET("/v1/hello", helloHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
