package main

import (
	"Webgame/server/api"
	"Webgame/server/serve"
	"Webgame/server/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
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
	api.CreateAPI(router) // API first because else site for api will get loaded TODO change port
	serve.CreateServe(router)

	util.Log(util.MAIN, "Startup complete")

	server := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().Port), Handler: router}
	util.Log(util.MAIN, fmt.Sprintf("ListenAndServe started on localhost%s", server.Addr))
	server.ErrorLog = log.New(util.LogWriter{}, "", 0)

	// blocks if success
	err = server.ListenAndServe()
	// TODO ListenAndServeTLS

	if err != nil {
		util.Err(util.MAIN, err, true, "Error serving site")
	}
}
