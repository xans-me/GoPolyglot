package middleware

import (
	"bytes"
	"net/http"
)

// bufferedResponseWriter wraps a http.ResponseWriter, capturing all data written to it.
type bufferedResponseWriter struct {
	http.ResponseWriter               // Embeds the original ResponseWriter to ensure all methods are available
	body                *bytes.Buffer // Buffer to capture the response
	status              int           // HTTP status code for the response
}

// newBufferedResponseWriter creates a new instance of bufferedResponseWriter.
func newBufferedResponseWriter(w http.ResponseWriter) *bufferedResponseWriter {
	return &bufferedResponseWriter{
		ResponseWriter: w,
		body:           new(bytes.Buffer),
		status:         http.StatusOK, // Default status code
	}
}

// Write captures data written to the response and writes it to the internal buffer.
func (b *bufferedResponseWriter) Write(data []byte) (int, error) {
	b.body.Write(data)                  // Capture data in the buffer
	return b.ResponseWriter.Write(data) // Write data to the original ResponseWriter
}

// WriteHeader captures the HTTP status code set for the response.
func (b *bufferedResponseWriter) WriteHeader(statusCode int) {
	b.status = statusCode                    // Capture the status code
	b.ResponseWriter.WriteHeader(statusCode) // Set the status code to the original ResponseWriter
}

// BodyString returns the contents of the buffer as a string.
func (b *bufferedResponseWriter) BodyString() string {
	return b.body.String()
}

// Status returns the captured HTTP status code.
func (b *bufferedResponseWriter) Status() int {
	return b.status
}
