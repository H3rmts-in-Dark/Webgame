package main

import (
	"fmt"
	"net/http"

	"Webgame/serve"
	"Webgame/util"

	"github.com/gorilla/mux"
)

func main() {
	util.LoadConfig()
	serve.LoadSite()
	util.Log("MAIN", "Starting Server...")
	util.Log("MAIN", "Startup complete")

	router := mux.NewRouter().StrictSlash(true)
	serve.CreateServe(router)

	err := http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Log("MAIN", "Error: ", err)
}
