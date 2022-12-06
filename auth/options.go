package auth

type Option func(a *Auth)

func WithUsername(username string) Option {
	return func(a *Auth) {
		a.Username = username
	}
}

func WithPassword(password string) Option {
	return func(a *Auth) {
		a.Password = password
	}
}
