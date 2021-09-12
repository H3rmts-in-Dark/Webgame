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
	
	util.Log(util.MAIN, "Starting Server")
	
	err = serve.Loadsite()
	if err != nil {
		return
	}
	util.Log(util.MAIN, "Loaded Sites")
	
	router := mux.NewRouter().StrictSlash(true)
	serve.CreateServe(router)
	api.CreateAPI(router)
	
	util.Log(util.MAIN, "Startup complete")
	
	// bocks if success
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Err(util.MAIN, err, true,"Error serving site")
}
