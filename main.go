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

	if util.GetConfig().Cache {
		err = serve.LoadSites()
		if err != nil {
			return
		}
	}

	webRouter := mux.NewRouter().StrictSlash(true)
	serve.CreateServe(webRouter)
	webServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().Port), Handler: webRouter}
	webServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.SERVER}, "", 0)

	apiRouter := mux.NewRouter().StrictSlash(true)
	api.CreateAPI(apiRouter)
	apiServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().ApiPort), Handler: apiRouter}
	apiServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.SERVERAPI}, "", 0)

	util.Log(util.MAIN, "Startup complete")

	go func() {
		util.Log(util.MAIN, fmt.Sprintf("ListenAndServe API started on api%s", apiServer.Addr))
		// blocks if success
		err = apiServer.ListenAndServe()
		// TODO ListenAndServeTLS

		if err != nil {
			util.Err(util.MAIN, err, true, "Error starting api")
		}
	}()

	util.Log(util.MAIN, fmt.Sprintf("ListenAndServe Webserver started on localhost%s", webServer.Addr))
	// blocks if success
	err = webServer.ListenAndServe()
	// TODO ListenAndServeTLS

	if err != nil {
		util.Err(util.MAIN, err, true, "Error starting webServer")
	}
}
