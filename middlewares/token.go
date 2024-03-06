package middlewares

import "net/http"

func ValidateExpAccessToken() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		funcHttp := func(w http.ResponseWriter, r *http.Request) {

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(funcHttp)
	}
}
