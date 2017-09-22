package model

import (
	"time"
)

const connection_timeout int = 0

type Connection struct{
	Uid string `json:"uid"`
	DeviceType string `json:"deviceType"`
	UserName string `json:"userName"`
	MobileAlive bool `json"mobileAlive"`
	BrowserAlive bool	`json":browserAlive"`
	LastTimeBrowserSync time.Time `json:"lastTimeBrowserSync"`
	LastTimeMobileSync	time.Time `json:lastTimeMobileSync`
	ForcedConnect bool `json:forcedConnect"`
}

//func ( conn *Connection) SetBrowserAlive() {
//	fmt.Printf("The browser is active")
//	fmt.Printf("The name of the device type is %s",conn.DeviceType)
//	conn.BrowserAlive = true
//	conn.LastTimeBrowserSync = time.Now()
//}
//
//func (conn *Connection) SetMobileAlive(){
//	conn.MobileAlive = true
//	conn.LastTimeMobileSync = time.Now()
//}
//
//func (conn * Connection)GetConnection( client *redis.Client, uid string) (*Connection,error){
//
//	value,err := client.Get(uid).Result()
//	if err != nil{
//		return nil,err
//	}
//	var connection *Connection
//	json.Unmarshal([]byte(value),connection)
//	return connection,nil
//
//}

//func (conn *Connection) SetConnection(db *redis.Client)(*Connection,error){
//
//	serialized,err := json.Marshal(conn)
//	if err != nil{
//		return nil,err
//	}
//	err = db.Set(conn.Uid,serialized,0).Err()
//	if err != nil{
//		return nil,err
//	}
//	return conn,nil
//
//}



func (conn *Connection)Update(connectNew *Connection)(*Connection,error){

	if connectNew.Uid !=""{
		conn.Uid= connectNew.Uid
	}
	if connectNew.DeviceType !=""{
		conn.DeviceType = connectNew.DeviceType
	}
	if connectNew.UserName !=""{
		conn.UserName = connectNew.UserName
	}
	if connectNew.MobileAlive{
		conn.MobileAlive = connectNew.MobileAlive
	}
	if connectNew.BrowserAlive{
		conn.BrowserAlive  = connectNew.BrowserAlive
	}
	if connectNew.DeviceType=="browser"{
		conn.LastTimeBrowserSync = time.Now()
	}
	if connectNew.DeviceType=="mobile"{
		conn.LastTimeMobileSync = time.Now()
	}
	return conn,nil
}
