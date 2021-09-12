package main

import (
	"fmt"
	"net/http"
	"os"
	
	"Webgame/api"
	"Webgame/serve"
	"Webgame/util"
	
	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("\nStarting Webgame Server Program args:%s\n\n", os.Args[1:])
	
	err := util.LoadConfig()
	if err != nil {
		return
	}
	
	util.Log("MAIN", "Starting Server")
	
	err = serve.Loadsite()
	if err != nil {
		return
	}
	util.Log("MAIN", "Loaded Sites")
	
	router := mux.NewRouter().StrictSlash(true)
	serve.CreateServe(router)
	api.CreateAPI(router)
	
	util.Log("MAIN", "Startup complete")
	
	// bocks if success
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Log("MAIN", "Error: ", err)
}
