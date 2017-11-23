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
	print("i ammmmm in chat handler"+ input)
	fmt.Fprintf(w, output)
	//fmt.Println(output)
}

func main() {

	dir := http.Dir("./static")
	fileServer := http.FileServer(dir)

	http.Handle("/", fileServer)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)

	http.ListenAndServe(":8080", nil)


}