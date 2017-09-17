package main

import (
	"github.com/facebookgo/inject"
	"github.com/gorilla/mux"
	"github.com/r-a-x/mAuth/repository"
	"github.com/r-a-x/mAuth/controller"
	"net/http"
	"github.com/go-redis/redis"
	"fmt"
	"os"
)


type App struct {
	ConnectionController *controller.ConnectionController `inject:""`
}

func initDB()(*redis.Client){

	client  :=  redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Printf(client.Ping().Result())
	return client
}

func DI()(*mux.Router){

	router := mux.NewRouter()
	var app App

	var g inject.Graph

	//Add the DI for the method
	err := g.Provide(
		&inject.Object{Value:&app},
		&inject.Object{Value:router},
		&inject.Object{Value:initDB()},
		&inject.Object{Value:controller.ConnectionControllerDI()},
		&inject.Object{Value:repository.ConnectionRepositoryDI()},
	)
	if err!=nil{
		panic(err)
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//Add the init call for the controller
	app.ConnectionController.Init()

	if err != nil{
		panic("Error Doing DI")
	}

	return router

}

func main() {

	router := DI()
	err := http.ListenAndServe("localhost:8080",router)
	if err != nil{
		panic(err)
	}

}
