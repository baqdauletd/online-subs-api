package router

import (
	"net/http"
	"online-subs-api/handlers"
)

func Routes(mux *http.ServeMux, subsHandler *handlers.SubsHandler){
	mux.HandleFunc("/subs/create", subsHandler.CreateSubHandler)
	mux.HandleFunc("/subs/getById", subsHandler.GetSubHandlerByID)
	mux.HandleFunc("/subs/listAll", subsHandler.ListAllSubsHandler)
	mux.HandleFunc("/subs/update", subsHandler.UpdateSubHandler)
	mux.HandleFunc("/subs/delete", subsHandler.DeleteSubHandler)
}