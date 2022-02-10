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
)

func main() {
	err := util.LoadConfig()
	if err != nil {
		panic(err)
	}

	util.Log(util.MAIN, "Starting server")
	logging.SQLInit()

	if util.GetConfig().Cache {
		err = serve.LoadSites()
		if err != nil {
			panic(err)
		}
	}

	serv := serve.CreateServe()

	if util.GetConfig().EnableHTTP {
		webServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().PortHTTP), Handler: serv}
		webServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.SERVER}, "", 0)
		go startwebServer(webServer, false)
	}
	if util.GetConfig().EnableHTTPS {
		webServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().PortHTTPS), Handler: serv}
		webServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.SERVER}, "", 0)
		go startwebServer(webServer, true)
	}

	http.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: &graph.Resolver{}})))
	APIServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().ApiPort), Handler: http.DefaultServeMux}
	APIServer.ErrorLog = log.New(&util.LogWriter{Prefix: util.GRAPHQL}, "", 0)

	startAPI(APIServer, util.GetConfig().ApiHTTPS)
}

func startwebServer(webServer *http.Server, tls bool) {
	// blocks if success
	var err error
	if tls {
		util.Log(util.MAIN, fmt.Sprintf("ListenAndServe Webserver with TLS started on localhost%s", webServer.Addr))
		err = webServer.ListenAndServeTLS(util.CertsFile, util.KeyFile)
	} else {
		util.Log(util.MAIN, fmt.Sprintf("ListenAndServe Webserver started on localhost%s", webServer.Addr))
		err = webServer.ListenAndServe()
	}

	if err != nil {
		util.Err(util.MAIN, err, true, "Error starting webServer")
	}
}

func startAPI(api *http.Server, tls bool) {
	// blocks if success
	var err error
	if tls {
		util.Log(util.MAIN, fmt.Sprintf("ListenAndServe API with TLS started on localhost%s", api.Addr))
		err = api.ListenAndServeTLS(util.CertsFile, util.KeyFile)
	} else {
		util.Log(util.MAIN, fmt.Sprintf("ListenAndServe API started on localhost%s", api.Addr))
		err = api.ListenAndServe()
	}

	if err != nil {
		util.Err(util.MAIN, err, true, "Error starting Api")
	}
}
