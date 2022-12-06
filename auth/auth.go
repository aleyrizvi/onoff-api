package auth

import (
	"github.com/labstack/echo/v4"
)

/* TODO: I am working on echo authentication framework to overcome the limitation of echo default middlewares.
   In future, this code is going to be replaced by auth package.
*/

type Auth struct{}

func New() *Auth {
	return &Auth{}
}

func (a *Auth) BasicAuth(username string, password string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			u, p, ok := c.Request().BasicAuth()

			// in case of Authorization header is invalid
			if !ok {
				return echo.ErrUnauthorized
			}

			// verify credentials
			if username != u || password != p {
				return echo.ErrUnauthorized
			}

			return nil
		}
	}
}
