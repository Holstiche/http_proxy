//This is executable code so package is main
package main

//all needed packages are imported here
import (
	"fmt"
	"net/http"
	"log"
)

//Here our handler function is definded
func handler(w http.ResponseWriter, r *http.Request) {
	//The url of the server the client wants to acces via the proxy is passed as a url-argument
	url := r.URL.Path[1:]
	//wordt gebruikt voor debuggen
	fmt.Fprintf(w, url)	
	//url validatie en flexibiliteit moet nog geimplementeerd worden

	if r.Method == "GET" {
		fmt.Fprintf(w, "r.Method is GET")
		resp, err := http.Get(url)
		//error handling
		if err != nil {
			log.Fatal(err)
		}
		//is dit wel echt nodig???
		defer resp.Body.Close()
	}
}


func main() {
	//The handler function is bound to the "/" url
	http.HandleFunc("/", handler)
	//The server is created
	log.Fatal(http.ListenAndServe(":8080", nil))
}