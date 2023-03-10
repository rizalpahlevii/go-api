package exception

type LoginError struct {
	Error string
}

func NewLoginError(error string) LoginError {
	return LoginError{
		Error: error,
	}
}
