package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hey there</h1>
<p>This is a multi-line string</p>
<p>in one Fprintf</p>
`)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}

/* HOW TO NOT BOTHER FIREWALL
-> dont use 	$: go run
-> use			$: go build -o a.exe && a.exe
*/
