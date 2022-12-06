package main

import (
	"fmt"
	"net/http"

	"github.com/aleyrizvi/onoff-api/anagram"
	auth "github.com/aleyrizvi/onoff-api/auth"
	"github.com/aleyrizvi/onoff-api/cmd/handlers"
	"github.com/aleyrizvi/onoff-api/validate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRouter(port int) {
	// echo router
	e := echo.New()

	// validator
	v := validate.NewValidatorInstance()
	e.Validator = v

	// services
	anagramService := anagram.New()

	// custom middlewares
	authenticator := auth.New(
		auth.WithUsername("admin"),
		auth.WithPassword("password"),
	)

	// generic middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(
		middleware.RequestID(),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
		//middleware.CORSWithConfig(middleware.DefaultCORSConfig),
		middleware.BodyLimit("2M"),
		authenticator.BasicAuth,
	)

	// handlers
	a := handlers.NewAnagramHandler(anagramService)

	// routes
	e.POST("/check-anagram", a.HandleCheckAnagram)

	// Start echo server
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}
