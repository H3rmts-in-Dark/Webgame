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
)

func main() {
	util.LoadConfig()
	util.ValidateConfig()
	logging.Log(logging.CONFIG, "Loaded config:", fmt.Sprintf("%+v", util.GetConfig()))

	logging.Log(logging.MAIN, "Starting server")
	logging.DBInit()

	if util.GetConfig().Cache {
		err := serve.LoadSites()
		if err != nil {
			panic(err)
		}
	}

	serv := serve.CreateServe()

	if util.GetConfig().EnableHTTP {
		webServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().PortHTTP), Handler: serv}
		webServer.ErrorLog = log.New(&logging.LogWriter{Prefix: logging.SERVER}, "", 0)
		go startWebServer(webServer, false)
	}
	if util.GetConfig().EnableHTTPS {
		webServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().PortHTTPS), Handler: serv}
		webServer.ErrorLog = log.New(&logging.LogWriter{Prefix: logging.SERVER}, "", 0)
		go startWebServer(webServer, true)
	}

	http.HandleFunc("/", graph.GetPlayground)
	http.Handle("/query", handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: graph.GenResolver()})))
	APIServer := &http.Server{Addr: ":" + fmt.Sprintf("%d", util.GetConfig().ApiPort), Handler: http.DefaultServeMux}
	APIServer.ErrorLog = log.New(&logging.LogWriter{Prefix: logging.GRAPHQL}, "", 0)

	startAPI(APIServer, util.GetConfig().ApiHTTPS)
}

func startWebServer(webServer *http.Server, tls bool) {
	// blocks if success
	var err error
	if tls {
		logging.Log(logging.MAIN, fmt.Sprintf("ListenAndServe Webserver with TLS started on localhost%s", webServer.Addr))
		err = webServer.ListenAndServeTLS(util.CertsFile, util.KeyFile)
	} else {
		logging.Log(logging.MAIN, fmt.Sprintf("ListenAndServe Webserver started on localhost%s", webServer.Addr))
		err = webServer.ListenAndServe()
	}

	if err != nil {
		logging.Err(logging.MAIN, err, "Error starting webServer")
	}
}

func startAPI(api *http.Server, tls bool) {
	// blocks if success
	var err error
	if tls {
		logging.Log(logging.MAIN, fmt.Sprintf("ListenAndServe API with TLS started on localhost%s", api.Addr))
		err = api.ListenAndServeTLS(util.CertsFile, util.KeyFile)
	} else {
		logging.Log(logging.MAIN, fmt.Sprintf("ListenAndServe API started on localhost%s", api.Addr))
		err = api.ListenAndServe()
	}

	if err != nil {
		logging.Err(logging.MAIN, err, "Error starting Api")
	}
}
