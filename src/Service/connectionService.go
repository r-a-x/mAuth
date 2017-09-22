package Service

import (
	"github.com/r-a-x/mAuth/src/model"
	"github.com/r-a-x/mAuth/src/repository"
	"github.com/r-a-x/mAuth/src/request"
	"time"
	"net/http"
	"fmt"
	"errors"
)

const expiryTime float64 = 100

var StreamNotSupportedError = errors.New("Stream not supported")

type ConnectionService struct {
	ConnectionRepository *repository.ConnectionRepository `inject:""`
	BrokerService  * Broker `inject:""`
}

func (service *ConnectionService) isConnectionExpired(timeStamp time.Time)(bool){

	var timeNull time.Time

	if timeStamp == timeNull{
		return true
	}

	return time.Since(timeStamp).Seconds() > expiryTime

}

func (service *ConnectionService) isConnectionRequired(isForceConnect bool,
	connectTo string, connection *model.Connection)(bool){

		if isForceConnect || connection == nil{
			return true
		}

		if connectTo == "mobile"{
			if connection.MobileAlive{
				return service.isConnectionExpired( connection.LastTimeMobileSync)
			}
		}

		if connectTo == "browser"{
			if connection.BrowserAlive{
				return service.isConnectionExpired( connection.LastTimeBrowserSync)
			}
		}

	return true

}

func (service *ConnectionService) Connect(connectionRequest *request.CreateConnectionRequest,w http.ResponseWriter)(*model.Connection,error) {

	connDb,err := service.ConnectionRepository.GetConnection(connectionRequest.Uid)

	if err !=nil{
		panic("Error Reading the value from the database")
	}

	if service.isConnectionRequired(connectionRequest.ForceConnect,
		connectionRequest.ConnectTo,connDb){
			return service.connect(connectionRequest,w)
	}

	return connDb,nil

}


func (service *ConnectionService) connect(connectionRequest *request.CreateConnectionRequest,w http.ResponseWriter)(*model.Connection,error){
	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return nil,StreamNotSupportedError
	}


	connectionRequest.ConnectStatus = make(chan model.ConnectStatus )

	service.BrokerService.Request <- *connectionRequest

	notify := w.(http.CloseNotifier).CloseNotify()

	go func() {

		<-notify
		// Remove this client from the map of attached clients
		// when `EventHandler` exits.
		//b.defunctClients <- messageChan

		close(connectionRequest.ConnectStatus)
		fmt.Println("HTTP connection just closed.")

	}()


	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")


	for {

		// Read from our messageChan.
		msg, open := <- connectionRequest.ConnectStatus

		if !open {
			fmt.Println("The connection is closed, Please try again")
			break
		}
		//
		//// Write to the ResponseWriter, `w`.
		fmt.Fprintf(w, "data: Message: %s\n\n", msg)
		//
		// Flush the response.  This is only possible if
		// the repsonse supports streaming.

		time.Sleep(200*time.Millisecond)
		f.Flush()

		close(connectionRequest.ConnectStatus)

	}

// Add a return statement
	return nil,nil
}

//This will test if the Browser or Mobile is connected to the backend or not
//If the other device is also available it will also return the status of that device
func (service *ConnectionService) IsConnected(connection *model.Connection) (*model.Connection,error){

	connDb,err:=service.ConnectionRepository.GetConnection(connection.Uid)

	if connection ==nil{
		connDb =  new(model.Connection)
	}

	if err!=nil{
		panic("Error Reading the value from the database")
	}

	if connection.DeviceType =="mobile"{
		connDb.MobileAlive = true
	}else{
		connDb.BrowserAlive = true
	}

	connDb, _ = connDb.Update(connection)

	service.ConnectionRepository.SetConnection(connDb)

	return connDb,nil
}
