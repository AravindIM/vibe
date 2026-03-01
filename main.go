package main

import (
	"sync"

	webview "github.com/AravindIM/webview_go"
)

const address = "http://localhost:8080/"

func web(wg *sync.WaitGroup) {
	defer wg.Done()
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Vibe")
	w.SetSize(480, 320, webview.HintNone)
	w.Navigate(address)
	w.Run()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go web(&wg)
	wg.Wait()
}
