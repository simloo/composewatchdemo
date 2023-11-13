package composewatchdemo

import (
	"fmt"
	"net/http"
)

func Serve() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "1.0.6")
	})

	http.ListenAndServe(":8080", nil)
}
