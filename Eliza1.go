//David Clarke G00335563
package main

//imports
import (
    //"html/template"
    "fmt"
    "net/http"
	"./regexp"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("inputUser")
	output := hello.ElizaOutput(input)
	//output:= "Hello"
	fmt.Fprintf(w, output)
	//fmt.Println(output)
}

func main() {

	dir := http.Dir("./static")
	Server := http.FileServer(dir)

	http.Handle("/", Server)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)

	http.ListenAndServe(":8080", nil)


}