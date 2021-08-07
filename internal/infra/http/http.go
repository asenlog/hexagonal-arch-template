package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// ReadTimeout HTTP Timeout
	ReadTimeout = 5 * time.Second
	// WriteTimeout Write HTTP Timeout
	WriteTimeout = 5 * time.Second
)

// HTTP is the HTTP handler
type HTTP struct {
}

// New create a new HTTP handler
func New() *HTTP {
	return &HTTP{}
}

// routeBuilder returns a new router with all registered routes
func (h *HTTP) routeBuilder() http.Handler {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, nil) })
	r.GET("/metrics", func(c *gin.Context) { c.JSON(http.StatusOK, nil) })

	return r
}

// Server returns the configured http.Server
func (h *HTTP) Server() *http.Server {
	return &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      h.routeBuilder(),
		ReadTimeout:  ReadTimeout,
		WriteTimeout: WriteTimeout,
	}
}
