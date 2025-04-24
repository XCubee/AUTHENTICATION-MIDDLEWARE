
package main

import (
	"fmt"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "12345678" {
			http.Error(w, "Unauthorized Attempt ", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome X.Cube"))
}

func main() {
	http.Handle("/", authMiddleware(http.HandlerFunc(homeHandler)))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
