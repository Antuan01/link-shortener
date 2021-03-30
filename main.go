package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.LoadHTMLGlob("templates/*")

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.GET("/index", func(c *gin.Context) {
   c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})   
  })
  r.Run()
}
