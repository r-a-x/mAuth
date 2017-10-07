package controller

import (
	"github.com/r-a-x/mAuth/src/Service"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/r-a-x/mAuth/src/model"
	"encoding/json"
	"github.com/r-a-x/mAuth/src/request"
	"fmt"
)

func ConnectionControllerDI()(*ConnectionController){
	connectionController := ConnectionController{}
	return &connectionController
}

// Remember to add the Path over here
func (connectionController *ConnectionController) Init(){
	connectionController.Router.HandleFunc("/isconnected",connectionController.isConnected).Methods("POST")
	connectionController.Router.HandleFunc("/connect",connectionController.Connect).Methods("POST")
	connectionController.Router.HandleFunc("/creds",connectionController.getCreds).Methods("POST")
}

type ConnectionController struct {
	ConnectionService *Service.ConnectionService `inject:""`
	Router *mux.Router `inject:""`
}


func (ConnectionController *ConnectionController) Connect(w http.ResponseWriter, r *http.Request){

	connectionRequest := new(request.CreateConnectionRequest)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&connectionRequest); err != nil{
		panic("Error parsing the Body !!!")
	}

	ConnectionController.ConnectionService.Connect(connectionRequest,w)

}

func( connectionController * ConnectionController) isConnected(w http.ResponseWriter , r *http.Request){

	checkConnectionReqeuest:=new(request.CheckConnectionRequest)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&checkConnectionReqeuest); err != nil{
		panic("Error parsing the Body !!!")
	}

	connection := new(model.Connection)
	connection.Uid = checkConnectionReqeuest.Uid
	connection.DeviceType = checkConnectionReqeuest.DeviceType

	connection,_ = connectionController.ConnectionService.IsConnected(connection)

	bytes,_:=json.Marshal(connection)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w,string(bytes))

}


func ( connectionController* ConnectionController ) getCreds(w http.ResponseWriter,r *http.Request) {

	credsRequest := new(request.CredsRequest)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&credsRequest); err != nil {
		panic("Error parsing the Body !!!")
	}
	//fmt.Print("The Cred request is ")
	//fmt.Print(credsRequest)

	credentials := new(model.Cred)
	credentials.Uid = credsRequest.Uid
	credentials.Url = credsRequest.Url
	credentials.SubmitUrl = credsRequest.SubmitUrl

	creds, _ := connectionController.ConnectionService.GetCreds(credentials)

	//fmt.Print("The creds recevied after processing the request is")
	//fmt.Print(creds)
	bytes, _ := json.Marshal(creds)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(bytes))

}

//This needs to sent to confirmed with phone, than only sent to the phone for confirmation
// Need to send the creds to the phone, and after the phone confirmation the data will be routed back

func( ConnectionController* ConnectionController) setCred(w http.ResponseWriter, r *http.Request){

	cred := new (model.Cred)
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&cred); err != nil {
		panic("Error parsing the Body !!!")
	}

	err := ConnectionController.ConnectionService.SetCred(cred)

	if err != nil{
		fmt.Fprintln(w,err)
	}

}