package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// response, error
	response, _ := http.Get("https://jsonplaceholder.typicode.com/todos")
	bytes, _ := ioutil.ReadAll(response.Body)
	stringBody := string(bytes)
	response.Body.Close()

	fmt.Println(stringBody)
}
