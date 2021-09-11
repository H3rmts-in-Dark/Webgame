package main

import (
	"fmt"
	"net/http"
	
	"Webgame/util"
	
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("\nStarting Server...\n ")
	util.LoadConfig()
	
	router := mux.NewRouter().StrictSlash(true)
	
	err := http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Log("MAIN", "Err: ", err)
}
