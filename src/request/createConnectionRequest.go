package request

import (
	"github.com/r-a-x/mAuth/src/model"
)

type CreateConnectionRequest struct {
	Uid string `json:"uid"`
	DeviceType string `json:"deviceType"`
	ConnectTo string `json:"connectTo"`
	ForceConnect bool `json:"forceConnect"`
	ConnectStatus chan model.ConnectStatus
}