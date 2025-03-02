package main

import (
	"just"
	"net/http"
)

func main() {
	r := just.New()
	r.GET("/", func(c *just.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Just</h1>")
	})

	r.GET("/hello", func(c *just.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *just.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *just.Context) {
		c.JSON(http.StatusOK, just.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
