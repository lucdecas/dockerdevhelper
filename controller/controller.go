package controller

import (
	"fmt"
    "log"
	"net/http"
	"github.com/lucdecas/dockerdevhelper/cmdmanager"
)

var(
	listContainerPath = "/list"
	searchImagePath = "/search"
	startContainerPath = "/start"
	stopContainerPath = "/stop"
	resetContainerPath = "/reset"
)

func Init(){
	fmt.Println("Initializing controller ...")
	main()
}

func main()  {
	http.HandleFunc(listContainerPath, list)
	http.HandleFunc(searchImagePath, search)
	http.HandleFunc(startContainerPath, start)
	http.HandleFunc(stopContainerPath, stop)
	http.HandleFunc(resetContainerPath, reset)

    if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list Command was called")
	cmdmanager.List()
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