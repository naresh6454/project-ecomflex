package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// bodyLogWriter is a custom response writer that captures the response body
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write captures the response body
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLogger logs HTTP requests and responses
func RequestLogger(logger loggerPkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		
		// Log request
		path := c.Request.URL.Path
		method := c.Request.Method
		
		// Read and restore the request body for logging
		var reqBody []byte
		if c.Request.Body != nil {
			reqBody, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}
		
		// Create a custom response writer to capture the response body
		bodyLogWriter := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBuffer(nil),
		}
		c.Writer = bodyLogWriter
		
		// Process request
		c.Next()
		
		// Log response
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		
		// Log request and response
		logger.Info(
			"Request processed",
			"path", path,
			"method", method,
			"status", statusCode,
			"latency", latency.String(),
			"client_ip", c.ClientIP(),
		)
		
		// Log request and response bodies in debug level
		if logger != nil && statusCode != 200 && statusCode != 204 {
			// Only log bodies for non-success responses or in debug mode
			if len(reqBody) > 0 {
				logger.Debug("Request body", "body", string(reqBody))
			}
			
			if bodyLogWriter.body.Len() > 0 {
				logger.Debug("Response body", "body", bodyLogWriter.body.String())
			}
		}
	}
}