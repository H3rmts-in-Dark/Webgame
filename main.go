package main

import (
	"fmt"
	"net/http"
	
	"Webgame/serve"
	"Webgame/util"
	
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("\nStarting Server...\n ")
	util.LoadConfig()
	serve.Loadsite()
	
	router := mux.NewRouter().StrictSlash(true)
	serve.CreateServe(router)
	
	err := http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Log("MAIN", "Err: ", err)
}
