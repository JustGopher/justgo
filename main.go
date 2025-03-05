package main

import (
	"just"
	"net/http"
)

func main() {
	r := just.New()
	r.Use(just.Logger(), just.Recovery())
	r.GET("/", func(c *just.Context) {
		c.String(http.StatusOK, "Hello just\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *just.Context) {
		names := []string{"just"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
