package GoPolyglot

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GinWrapper wraps an agnostic http.Handler middleware for use with Gin.
func GinWrapper(mw func(http.Handler) http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a http.Handler from the Gin context
		handler := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		}))
		// Use the wrapped handler to serve the request
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
