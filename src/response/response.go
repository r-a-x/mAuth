package response

import (
	"time"
)

// Remember to update the time limit over here, only for when the last call to be made for the new Mobile etc
type CheckConnectionResponse struct{
	Uid string `json:"uid"`
	deviceAlive bool `json:"deviceAlive"`
	deviceUserName string `json:"deviceUserName"`
	LastTimeDeviceSync time.Time `json:"lastTimeDeviceSync"`
}

type CredsResponse struct{
	Uid string `json:"uid"`
	data string `json:"creds"`
}




//type createConnectionResponse struct {
//	Uid string `json:"uid"`,
//
//}