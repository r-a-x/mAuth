package controller

import (
	"github.com/r-a-x/mAuth/Service"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/r-a-x/mAuth/model"
	"encoding/json"
)

func ConnectionControllerDI()(*ConnectionController){
	connectionController := ConnectionController{}
	return &connectionController
}

// Remember to add the Path over here
func (connectionController *ConnectionController) Init(){
	connectionController.Router.HandleFunc("/isconnected",connectionController.isConnected).Methods("POST")
}

type ConnectionController struct {
	ConnectionService *Service.ConnectionService `inject:""`
	Router *mux.Router `inject:""`
}


func (ConnectionController *ConnectionController) Connect(w http.ResponseWriter, r *http.Request){
	connection := new(model.Connection)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&connection); err != nil{
		panic("Error parsing the Body !!!")
	}
	ConnectionController.ConnectionService.Connect(connection)
}

func( connectionController * ConnectionController) isConnected(w http.ResponseWriter , r *http.Request){
	connection := new(model.Connection)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&connection); err != nil{
		panic("Error parsing the Body !!!")
	}
	connectionController.ConnectionService.IsConnected(connection)
}




