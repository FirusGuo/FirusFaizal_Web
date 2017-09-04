package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//New httprouter
	router := httprouter.New()

	//Serve Assets Folders
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/images/*filepath", http.Dir("images"))
	router.ServeFiles("/sitemap/*filepath", http.Dir("sitemap"))

	//Routes
	router.GET("/", Index)
	router.GET("/robots.txt", robots)

	//Local Server
	// fmt.Println("Listening on Local Server localhost:3000...")
	// panic(http.ListenAndServe(":3000", router))

	//Remote Server
	fmt.Println("Listening on Remote Server...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}
}

//Index     Handler: Index, Route: /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		fmt.Println("Error:405:Method is Not GET Method, Handler: Index, Route: / ")
		return
	}
	http.ServeFile(w, r, "index.html")
}

//robots    Handler: robots, Route: /robots.txt
func robots(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		fmt.Println("Error:405:Method is Not GET Method, Handler: robots, Route: /robots.txt ")
		return
	}
	http.ServeFile(w, r, "robots.txt")
}
