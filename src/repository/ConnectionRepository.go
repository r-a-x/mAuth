package repository

import (
	"github.com/go-redis/redis"
	"github.com/r-a-x/mAuth/src/model"
	"encoding/json"
	"fmt"
)




type ConnectionRepository struct {
	Db *redis.Client `inject:""`
}

func ConnectionRepositoryDI()(*ConnectionRepository){
	connectionRepository := ConnectionRepository{}
	return &connectionRepository
}

func (repo *ConnectionRepository) GetConnection(uid string)(*model.Connection,error){

	value,_ := repo.Db.Get(uid).Result()
	fmt.Printf("The value of the %s\n",value)
	var connection = new (model.Connection)
	json.Unmarshal([]byte(value),connection)
	return connection,nil


}

func (repo *ConnectionRepository) SetConnection(connection *model.Connection)(*model.Connection,error){

	serializedBytes, _ :=json.Marshal(connection)
	serializedString := string(serializedBytes)

	err := repo.Db.Set(connection.Uid,serializedString,0 ).Err()
	if err != nil{
		return nil,err
	}

	return connection,nil
}

