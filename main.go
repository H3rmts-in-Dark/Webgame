package main

import (
	"fmt"
	"log"
	"net/http"

	graph "Webgame/server/graphql"
	gen "Webgame/server/graphql/generated"
	"Webgame/server/serve"
	"Webgame/server/util"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

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

	APIRouter := mux.NewRouter().StrictSlash(true)
	APIRouter.Handle("/", playground.Handler("GraphQL playground", "/query"))
	APIRouter.Handle("/query", handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: &graph.Resolver{}})))
	APIServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().ApiPort), Handler: APIRouter}
	APIServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.GRAPHQL}, "", 0)

	util.Log(util.MAIN, "Startup complete")

	go startAPI(APIServer)
	startwebServer(webServer)
}

func startwebServer(webServer *http.Server) {
	util.Log(util.MAIN, fmt.Sprintf("ListenAndServe Webserver started on localhost%s", webServer.Addr))
	// blocks if success
	err := webServer.ListenAndServeTLS("server/server.pem", "server/server.key")

	if err != nil {
		util.Err(util.MAIN, err, true, "Error starting webServer")
	}
}

func startAPI(Api *http.Server) {
	util.Log(util.MAIN, fmt.Sprintf("ListenAndServe API started on localhost%s", Api.Addr))
	// blocks if success
	err := Api.ListenAndServeTLS("server/server.pem", "server/server.key")

	if err != nil {
		util.Err(util.MAIN, err, true, "Error starting Api")
	}
}
