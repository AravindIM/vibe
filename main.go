package main

import (
	"fmt"
	"net/http"

	webview "github.com/AravindIM/webview_go"
	"github.com/gin-gonic/gin"
)

const address = "localhost:8080"

func web(done chan<- error) {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Vibe")
	w.SetSize(480, 320, webview.HintNone)
	w.Navigate(fmt.Sprintf("http://%s", address))
	w.Run()
	done <- fmt.Errorf("Closed")
}

func api(done chan<- error) {
	count := 0
	router := gin.Default()
	router.LoadHTMLFiles("./template/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"count": fmt.Sprintf("%d", count)})
	})
	router.POST("/add", func(c *gin.Context) {
		count += 1
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf("%d", count)))
	})
	router.POST("/sub", func(c *gin.Context) {
		count -= 1
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf("%d", count)))
	})
	done <- router.Run(address)
}

func main() {
	done := make(chan error)
	go web(done)
	go api(done)
	err := <-done
	fmt.Println("Exiting:", err)
}
