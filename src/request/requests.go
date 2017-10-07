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



type CheckConnectionRequest struct {
	Uid        string `json:"uid"`
	DeviceType string `json:"deviceType"`
	Ping       string `json:"ping"`
}

/*

    private String uid;
    private String url;
    private String submitUrl;
    private String username;
    private String password;
    private String name;
    private String StringFields;

*/

type CredsRequest struct {
	Uid string `json:"uid"`
	Url string `json:"url"`
	SubmitUrl string `json:"submitUrl"`
}
