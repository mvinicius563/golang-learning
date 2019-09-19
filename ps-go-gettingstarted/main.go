package main

import (
	"net/http"

	"github.com/mvinicius563/golang-learning/ps-go-gettingstarted/controllers"
)

func main() {
	// type Message struct {
	// 	Name string
	// 	Body string
	// 	Time int64
	// }
	// m := Message{"Alice", "Hello", 1294706395881547000}

	// b, _ := json.Marshal(m)
	// fmt.Printf("%v\n", string(b))
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
