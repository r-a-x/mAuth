package util

import (
	"io"
	"encoding/json"
)

func Decode(r io.Reader,ob * interface{} )(*interface{},error){
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&ob); err != nil{
		return nil,err
	}
	return ob,nil
}