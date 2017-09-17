package Service

import (
	"github.com/r-a-x/mAuth/model"
	"github.com/r-a-x/mAuth/repository"
	"github.com/r-a-x/mAuth/request"
	"time"
)

const expiryTime float64 = 100

type ConnectionService struct {
	ConnectionRepository *repository.ConnectionRepository `inject:""`
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

func (service *ConnectionService) Connect(connectionRequest *request.CreateConnectionRequest)(*model.Connection,error) {

	connDb,err := service.ConnectionRepository.GetConnection(connectionRequest.Uid)

	if err !=nil{
		panic("Error Reading the value from the database")
	}

	if service.isConnectionRequired(connectionRequest.ForceConnect,
		connectionRequest.ConnectTo,connDb){
			return service.connect(connectionRequest)
	}

	return connDb,nil

}


func (service *ConnectionService) connect(connectionRequest *request.CreateConnectionRequest)(*model.Connection,error){
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
