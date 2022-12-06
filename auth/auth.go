package auth

import (
	"github.com/labstack/echo/v4"
)

/* TODO: I am working on echo authentication framework to overcome the limitation of echo default middlewares.
   In future, this code is going to be replaced by auth package.
*/

type Auth struct {
	Username string
	Password string
}

func New(opts ...Option) *Auth {
	auth := &Auth{
		Username: "admin",
		Password: "password",
	}

	for _, opt := range opts {
		opt(auth)
	}
	return auth
}

func (a *Auth) BasicAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, p, ok := c.Request().BasicAuth()

		// in case of Authorization header is invalid
		if !ok {
			return echo.ErrUnauthorized
		}

		// verify credentials
		if u != a.Username || p != a.Password {
			return echo.ErrUnauthorized
		}

		return nil
	}
}
