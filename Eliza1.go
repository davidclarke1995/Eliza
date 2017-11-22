//David Clarke G00335563
package main

//imports
import (
    //"html/template"
    "fmt"
    "net/http"
)

    func routeHandler(w http.ResponseWriter, r *http.Request) {
	//serve the homepage.html file
	http.ServeFile(w, r, "Eliza.html")
    w.Header().Set("Content-Type", "text/html")

}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	output:= "Hello"
	fmt.Fprintf(w, output)
}

func main() {


	http.HandleFunc("/", routeHandler)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)

	http.ListenAndServe(":8080", nil)


}