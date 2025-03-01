package main

import (
	"just"
	"net/http"
)

func main() {
	r := just.New()
	r.GET("/", func(c *just.Context) {
		c.HTML(http.StatusOK, "<h1>Hello just!</h1>")
	})

	r.GET("/hello", func(c *just.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("login", func(c *just.Context) {
		c.JSON(http.StatusOK, just.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
