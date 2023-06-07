package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Get the GIF file
	gif, err := http.Get("https://raw.githubusercontent.com/etychon/iox-multiarch-nginx-nyancat-sample/master/images/nyan.gif")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Serve the GIF file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", gif.Header.Get("Content-Type"))
		io.Copy(w, gif.Body)
	})

	http.ListenAndServe(":8080", nil)
}
