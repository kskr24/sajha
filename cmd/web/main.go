package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
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
	fmt.Printf("fuck off")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "invalid request"})
	}
	fmt.Print("Duh")
	mu.Lock()
	users = append(users, *u)
	mu.Unlock()

	return c.JSON(http.StatusOK,
		map[string]string{"message": "login successful"})
}

func listUserHandler(c echo.Context) error {
	mu.Lock()
	defer mu.Unlock()
	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"http://localhost:4200"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},},
	))
	e.POST("/v1/login", loginHandler)
	e.GET("/v1/users", listUserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
