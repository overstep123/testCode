package main

import (
	"fmt"

	kepperClient "github.com/qmessenger/gokeeper/client"
)

func main() {
	/*http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		return
	})
	http.ListenAndServe(":4444", nil)*/
	t := map[string]string{}
	kepperClient.New()
	fmt.Println(&t)
}
