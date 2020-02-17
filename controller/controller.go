package controller

import (
	"fmt"
    "log"
	"net/http"
	"github.com/lucdecas/dockerdevhelper/cmdmanager"
)

var(
	searchImageEndpoint = "/search"
	startContainerEndpoint = "/start"
	stopContainerEndpoint = "/stop"
	resetContainerEndpoint = "/reset"
)

func Init(){
	main()
}

func main() {
	http.HandleFunc(searchImageEndpoint, search)
	http.HandleFunc(startContainerEndpoint, start)
	http.HandleFunc(stopContainerEndpoint, stop)
	http.HandleFunc(resetContainerEndpoint, reset)

	
	
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}

func search(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "search Command was called")
	cmdmanager.Search(r.Header.Get("target"))
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "start Command was called")
	cmdmanager.Start(r.Header.Get("target"))
}

func stop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "stop Command was called")
	cmdmanager.Stop(r.Header.Get("target"))
}

func reset(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "reset Command was called")
	cmdmanager.Reset(r.Header.Get("target"))
}