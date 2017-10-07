package repository

import (
	"github.com/go-redis/redis"
	"github.com/r-a-x/mAuth/src/model"
	"encoding/json"
	"fmt"
)




type CredsRepository struct {
	Db *redis.Client `inject:""`
}

func CredsRepositoryDI()(*CredsRepository){
	credsRepository := CredsRepository{}
	return &credsRepository
}


func (repo *CredsRepository) GetCreds(uid string,url string, submitUrl string)([]model.Cred,error){

	value,_:=repo.Db.Get(uid+url+submitUrl).Result()
	fmt.Print("The value read from the redis database is ")
	fmt.Print(value)
	creds := make([]model.Cred,0)
	err := json.Unmarshal([]byte(value),&creds)
	if err != nil{
		fmt.Print("Some Error processing the db to array")
		fmt.Print(err)
	}
	return creds,nil
}

func (repo *CredsRepository) SetCreds(creds *model.Cred)(error){

	serializedBytes, _:= json.Marshal(creds)
	serializedString := string(serializedBytes)
	err := repo.Db.Set(creds.Uid+creds.Url+creds.SubmitUrl,serializedString,0).Err()
	return err
}

