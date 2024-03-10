package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// WebContext defines an interface for web context operations,
// making it easier to write middleware that works across different web frameworks.
type WebContext interface {
	JSON(statusCode int, obj interface{})
	SetHeader(key, value string)
	GetHeader(key string) string
	Next()
	Request() *http.Request
	Writer() http.ResponseWriter
}

// GinWebContext implements the WebContext interface using *gin.Context,
// allowing middleware to interact with Gin's context in a generic way.
type GinWebContext struct {
	*gin.Context
}

// JSON sends a JSON response with the given status code and object.
func (g *GinWebContext) JSON(statusCode int, obj interface{}) {
	g.Context.JSON(statusCode, obj)
}

// SetHeader sets a response header.
func (g *GinWebContext) SetHeader(key, value string) {
	g.Context.Header(key, value)
}

// GetHeader retrieves a request header's value.
func (g *GinWebContext) GetHeader(key string) string {
	return g.Context.GetHeader(key)
}

// Next calls the next handler in the middleware chain.
func (g *GinWebContext) Next() {
	g.Context.Next()
}

// Request returns the original *http.Request.
func (g *GinWebContext) Request() *http.Request {
	return g.Context.Request
}

// Writer returns the http.ResponseWriter to write the response.
func (g *GinWebContext) Writer() http.ResponseWriter {
	return g.Context.Writer
}

// MuxWebContext implements the WebContext interface using http.ResponseWriter and *http.Request,
// enabling generic middleware interaction with the standard net/http library used by Mux.
type MuxWebContext struct {
	RW  http.ResponseWriter
	Req *http.Request
}

// JSON encodes the given object as JSON and sends it with the specified status code.
func (m *MuxWebContext) JSON(statusCode int, obj interface{}) {
	m.RW.Header().Set("Content-Type", "application/json")
	m.RW.WriteHeader(statusCode)
	json.NewEncoder(m.RW).Encode(obj)
}

// SetHeader sets a response header.
func (m *MuxWebContext) SetHeader(key, value string) {
	m.RW.Header().Set(key, value)
}

// GetHeader retrieves a request header's value.
func (m *MuxWebContext) GetHeader(key string) string {
	return m.Req.Header.Get(key)
}

// Next can be left empty or used to manually trigger the next middleware in Mux, if necessary.
func (m *MuxWebContext) Next() {
	// Implementation may vary based on how middleware chaining is handled in Mux.
}

// Request returns the original *http.Request.
func (m *MuxWebContext) Request() *http.Request {
	return m.Req
}

// Writer returns the http.ResponseWriter to write the response.
func (m *MuxWebContext) Writer() http.ResponseWriter {
	return m.RW
}
