package main

import (
	"fmt"
	"log"
	"net/http"

	graph "Server/graphql"
	gen "Server/graphql/generated"
	"Server/logging"
	"Server/serve"
	"Server/util"
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
	logging.SQLInit()

	if util.GetConfig().Cache {
		err = serve.LoadSites()
		if err != nil {
			return
		}
	}

	webRouter := mux.NewRouter().StrictSlash(true)
	serve.CreateServe(webRouter)

	if util.GetConfig().EnableHTTP {
		webServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().PortHTTP), Handler: webRouter}
		webServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.SERVER}, "", 0)
		go startwebServer(webServer, false)
	}
	if util.GetConfig().EnableHTTPS {
		webServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().PortHTTPS), Handler: webRouter}
		webServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.SERVER}, "", 0)
		go startwebServer(webServer, true)
	}

	APIRouter := mux.NewRouter().StrictSlash(true)
	APIRouter.Handle("/", playground.Handler("GraphQL playground", "/query"))
	APIRouter.Handle("/query", handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: &graph.Resolver{}})))
	APIServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().ApiPort), Handler: APIRouter}
	APIServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.GRAPHQL}, "", 0)

	startAPI(APIServer, util.GetConfig().ApiHTTPS)
}

func startwebServer(webServer *http.Server, tls bool) {
	util.Log(util.MAIN, fmt.Sprintf("ListenAndServe Webserver started on localhost%s", webServer.Addr))
	// blocks if success
	var err error
	if tls {
		err = webServer.ListenAndServeTLS("server.pem", "server.key")
	} else {
		err = webServer.ListenAndServe()
	}

	if err != nil {
		util.Err(util.MAIN, err, true, "Error starting webServer")
	}
}

func startAPI(api *http.Server, tls bool) {
	util.Log(util.MAIN, fmt.Sprintf("ListenAndServe API started on localhost%s", api.Addr))
	// blocks if success
	var err error
	if tls {
		err = api.ListenAndServeTLS("server.pem", "server.key")
	} else {
		err = api.ListenAndServe()
	}

	if err != nil {
		util.Err(util.MAIN, err, true, "Error starting Api")
	}
}
