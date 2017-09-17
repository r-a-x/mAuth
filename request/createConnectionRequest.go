package request


type CreateConnectionRequest struct {
	Uid string `json:"uid"`
	DeviceType string `json:"deviceType"`
	ConnectTo string `json:"connectTo"`
	ForceConnect bool `json:"forceConnect"`
}