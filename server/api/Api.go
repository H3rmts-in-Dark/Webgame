package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"Webgame/server/util"

	"github.com/gorilla/mux"
)

type Action string

const (
	Invalid    Action = "Invalid"
	ReloadSite Action = "Reload_Site"
)

/*
checks if received JSON has APIkey and (Register or Activity or Log or Action) as key.
Returns the API key and Requesttype.
*/
func validateAPIJSON(js *map[string]interface{}) Action {
	action, actionExists := (*js)["action"]
	if actionExists {
		switch toAction(action) {
		case ReloadSite:
			return ReloadSite
		}
	}
	return Invalid
}

func toAction(action interface{}) Action {
	return Action(fmt.Sprintf("%v", action))
}

/*
This function is called when a Connection sends data.
It turns the raw byte array into a JSON to process.
Afterwards, it returns a response byte array made of JSON data to write back at the connection.
*/
func receiveAPI(raw *[]byte) []byte {
	fmt.Println()

	var receive map[string]interface{}
	err := json.Unmarshal(*raw, &receive)

	if err != nil {
		util.Err(util.API, err, false, "JSON decoding error:", string(*raw))
		return nil
	}
	if len(receive) == 0 {
		util.Err(util.API, nil, false, "Empty JSON request")
		return nil
	}
	action := validateAPIJSON(&receive)
	if action == Invalid {
		util.Err(util.API, nil, false, "Invalid request action", receive)
		return nil
	}
	util.Log(util.API, "received:", receive)

	var data interface{}

	switch action {
	case ReloadSite:
		err = CheckAdmin(&receive)
		if err == nil {
			data, err = ProcessSiteReloadRequest()
		}
	}

	if err != nil {
		util.Err(util.API, err, false, "Error processing request")
		msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": err.Error()})
		return msg
	} else {
		msg, err := json.Marshal(map[string]interface{}{"action": action, "data": data})
		if err != nil {
			util.Err(util.API, err, true, "Error sending answer")
			msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": err.Error()})
			return msg
		}
		return msg
	}
}

/*
CreateAPI

Registers '/api' handle to mux.Router with a JSON return with the POST method.
*/
func CreateAPI(rout *mux.Router) {
	rout.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := ioutil.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")

		msg := receiveAPI(&raw)

		if msg == nil {
			w.WriteHeader(http.StatusBadRequest)
			msg, _ = json.Marshal(map[string]interface{}{"error": "Bad request"})
		}

		_, err := w.Write(msg)
		if err != nil {
			util.Err(util.API, err, true, "Error writing response:")
		} else {
			util.Log(util.API, "sent:", string(msg))
		}
	}).Methods("POST")
}
