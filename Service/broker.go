package Service

import (
	"fmt"
	"github.com/r-a-x/mAuth/request"
	"github.com/r-a-x/mAuth/model"
)

type Broker struct{
	Request chan request.CreateConnectionRequest
}

func connectToDevice(to request.CreateConnectionRequest){

	// This method will make an actual request to the method and will return the result from there
	fmt.Println("In the ConnectToDevice Method and waiting for the response from the ")
	connectStatus  := model.ConnectStatus {}
	connectStatus.Uid = to.Uid
	connectStatus.DeviceConected = true
	connectStatus.DeviceType = to.DeviceType
	to.ConnectStatus <- connectStatus
	fmt.Println("Going to send something to the method")

}

func (b *Broker)Start(){

	go func(){
		for {

			fmt.Println("The Broker Started !!")
			select {
			case s := <-b.Request:
				fmt.Println("Got the new request to connect to the mobile or the backend")
				go connectToDevice(s)
			}

		}
	}()

}
