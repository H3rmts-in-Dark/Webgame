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
	fmt.Printf("Starting webgame server program with arguments: %s\n", os.Args[1:])

	err := util.LoadConfig()
	if err != nil {
		return
	}

	util.Log(util.MAIN, "Starting server")

	err = serve.Loadsites()
	if err != nil {
		return
	}

	router := mux.NewRouter().StrictSlash(true)
	serve.CreateServe(router)
	api.CreateAPI(router)

	util.Log(util.MAIN, "Startup complete")

	// blocks if success
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", util.GetConfig().Port), router)
	util.Err(util.MAIN, err, true, "Error serving site")
}
