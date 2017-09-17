package controller

import (
	"github.com/r-a-x/mAuth/Service"
	"github.com/gorilla/mux"
)

type ErrorController struct {
	ErrorService Service.ErrorService `inject:""`
	Router *mux.Router
}

func ErrorControllerDI()(*ErrorController){
	return new (ErrorController)
}

func ( ErrorController * ErrorController) init(){
	//ErrorController.Router.NotFoundHandler.ServeHTTP()
}

