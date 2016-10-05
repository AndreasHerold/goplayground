package main

import "fmt"

//import "net/http"
import "local/an.herold/page"

func main() {
	//	http.HandleFunc("/", handler)
	//	http.ListenAndServe(":8000", nil)
	handler()
}

func handler() {
	fmt.Println(page.SayHello())

}
