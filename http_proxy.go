//This is executable code so pakage is main
package main

//all needed packages are imported here
import (
	"fmt"
	"net/http"
)

//Here our handler function is definded
func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	fmt.Fprintf(w, url)
}

//our main code shloud come here
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}