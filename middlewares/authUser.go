package middlewares

import (
	"fmt"
	"net/http"
)

func TokenMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Your logic to retrieve the token from the request and send it to the third-party service
        token := r.Header.Get("Authorization")
		  if token == "" {
			
		  }
        // Send the token to the third-party service
        // For demonstration purposes, let's just print it
        fmt.Println("Token sent to third-party service:", token)

        // Call the next handler in the chain
        next.ServeHTTP(w, r)
    })
}
