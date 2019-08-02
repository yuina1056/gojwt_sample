package main

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/login", func(c echo.Context) error {
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "yuina"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
		if err != nil {
		}
		return c.String(http.StatusOK, tokenString)
	})
	r := e.Group("/auth")
	r.Use(middleware.JWT([]byte(os.Getenv("SIGNINGKEY"))))
	r.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Test World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
