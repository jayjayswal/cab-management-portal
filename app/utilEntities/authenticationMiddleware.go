package utilEntities

import (
	"log"
	"net/http"
)

// Define our struct
type AuthenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *AuthenticationMiddleware) Populate() {
	amw.tokenUsers = make(map[string]string)
	amw.tokenUsers["1a36591bceec49c832079e270d7e8b73"] = "user0"
}

// Middleware function, which will be called for each request
func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth-token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}