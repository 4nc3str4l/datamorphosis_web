package main


import (
    "github.com/gin-gonic/gin"
    "net/http"
	"path/filepath"
)


func main() {
	r := gin.Default()
	// Load the HTML Folders
	htmlPath := filepath.FromSlash("./templates/**.html")
	r.LoadHTMLGlob(htmlPath)

	// Load the static files
	cssPath := filepath.FromSlash("./static/css")
	r.Static("/css", cssPath)
	
	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}