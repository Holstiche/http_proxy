package main

//all needed packages are imported here
import (
	"fmt"
	"log"
	"net/http"
	"io"
)

//To do
//2: when an url is clicked it should also be accesed through the proxy
//3: url fixing
//4: add header support, contenttype
//5: add other methods 
//6: remove fmt debugginh
//7: implement go routines for better performance

type Proxy struct {
}

func NewProxy() *Proxy { return &Proxy{} }

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("proxy accessed")
	var resp *http.Response
	var err error

	switch r.Method {
	default:
		{
			fmt.Println("Cannot handle method ", r.Method)
			return
		}
	case "GET":
		{
			fmt.Println("getting")
			fmt.Println(r.URL.String()[1:])
			resp, err = http.Get(r.URL.String()[1:])
			r.Body.Close()
		}
	case "POST":
		{
			fmt.Println("posting")
			//contenttype moet worden aangepast
			resp, err = http.Post(r.URL.String()[1:], r.Header["Content-Type"][0], r.Body)
			r.Body.Close()
		}
	}

	// combined for GET/POST
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(w, resp.Body)
	//resp.Write(w)
	defer resp.Body.Close()
}

func main() {
	proxy := NewProxy()
	log.Fatal(http.ListenAndServe(":8080/", proxy))
}
