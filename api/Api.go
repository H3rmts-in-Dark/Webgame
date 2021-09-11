package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
	"Webgame/util"
	
	"github.com/gorilla/mux"
)

/*
checks if recived JSON has APIkey and (Register or Activity or Log or Action) as key
returns the API key and Requesttype
*/
func validateAPIJSON(js *map[string]interface{}) Action {
	action, action_exists := (*js)["action"]
	if action_exists {
		switch toAction(action) {
		case ReloadSite:
			return ReloadSite
		}
	}
	return Invalid
}

type Action string

const (
	Invalid    Action = "invalid"
	ReloadSite Action = "Reload Site"
)

func toAction(action interface{}) Action {
	return Action(fmt.Sprintf("%v", action))
}

/*
called when Connection send data;
gets byte array out of JSON
returns byte array out of JSON to write
*/
func reciveAPI(raw *[]byte) []byte {
	fmt.Println()
	
	var recive map[string]interface{}
	err := json.Unmarshal(*raw, &recive)
	
	if err != nil {
		util.Log("API", "JSON decoding error: ", err, " :", string(*raw))
		return nil
	}
	if len(recive) == 0 {
		util.Log("API", "empty JSON")
		return nil
	}
	action := validateAPIJSON(&recive)
	if action == Invalid {
		util.Log("Api", "invalid action", recive)
		return nil
	}
	util.Log("Api", "recived:", recive)
	
	var data interface{}
	
	switch action {
	case ReloadSite:
		err = Checkadmin(&recive)
		if err == nil {
			data, err = ProcessSiteReloadRequest()
		}
	}
	
	if err != nil {
		util.Log("Api", "send err:", err)
		msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": err.Error()})
		return msg
	} else {
		msg, err := json.Marshal(map[string]interface{}{"action": action, "data": data})
		if err != nil {
			util.Log("Api", "send err:", err)
			msg, _ := json.Marshal(map[string]interface{}{"action": action, "error": err.Error()})
			return msg
		}
		return msg
	}
}

/*
CreateAPI

Registers /api handle to mux.Router with json return and POST get
*/
func CreateAPI(rout *mux.Router) {
	rout.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		raw, _ := ioutil.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		
		msg := reciveAPI(&raw)
		
		if msg == nil {
			w.WriteHeader(http.StatusBadRequest)
			msg, _ = json.Marshal(map[string]interface{}{"error": "bad request"})
		}
		
		_, err := w.Write(msg)
		if err != nil {
			util.Log("API", "err in sending:", err)
		} else {
			util.Log("API", "send:", string(msg))
		}
	}).Methods("POST")
}
