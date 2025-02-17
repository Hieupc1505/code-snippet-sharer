package api

type OptFunc func(a *API)

func WithHTTPS(secure bool) OptFunc {
	return func(a *API) {
		a.secure = secure
	}
}

func WithPort(port int) OptFunc {
	return func(a *API) {
		a.port = port
	}
}
