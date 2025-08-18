package handlers

import (
	"encoding/json"
	"net/http"
	"online-subs-api/models"
	"online-subs-api/services"
	"online-subs-api/utils"
)

type JSONSubRequest struct {
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type SubsHandler struct{
	subsService *services.SubsService
}

func NewSubHandler(subsService *services.SubsService) *SubsHandler{
	return &SubsHandler{subsService: subsService}
}

func (h *SubsHandler) CreateSubHandler(w http.ResponseWriter, r *http.Request){
	utils.InfoLogger.Println("CreateSubHandler called")
	
	var req JSONSubRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorLogger.Printf("Failed to decode request body: %v\n", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	sub := &models.Sub{
		ServiceName: req.ServiceName,
		Price: req.Price,
		UserID: req.UserID,
	}

	if err := h.subsService.CreateService(sub, req.StartDate, req.EndDate); err != nil {
		utils.ErrorLogger.Printf("Failed to create subscription: %v\n", err)
		http.Error(w, "failed to create subscription", http.StatusBadRequest)
		return
	}

	utils.InfoLogger.Printf("Subscirpition created succesfully: %+v\n", sub)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}

func (h *SubsHandler) GetSubHandlerByID(w http.ResponseWriter, r *http.Request){
	utils.InfoLogger.Println("GetSubHandlerByID called")
	id := r.URL.Query().Get("id")
	if id == ""{
		utils.WarningLogger.Println("Missing id parameter in request")
		http.Error(w, "missing id paramter", http.StatusBadRequest)
		return
	}

	sub, err := h.subsService.GetServiceByID(id)
	if err != nil{
		utils.ErrorLogger.Printf("Subscription not found for id=%s: %v", id, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.InfoLogger.Printf("Subscription retrieved successfully: %+v", sub)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

func (h *SubsHandler) ListAllSubsHandler(w http.ResponseWriter, r *http.Request){
	utils.InfoLogger.Println("ListAllSubsHandler Called")
	var req JSONSubRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorLogger.Printf("Failed to decode request body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	subs, err := h.subsService.ListAllSubsService(req.UserID, req.ServiceName)
	if err != nil{
		utils.ErrorLogger.Printf("Failed to list subscriptions: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	utils.InfoLogger.Println("Subs Listed successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subs)
}

func (h *SubsHandler) UpdateSubHandler(w http.ResponseWriter, r *http.Request){
	utils.InfoLogger.Println("UpdateSubHandler Called")
	var req JSONSubRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorLogger.Printf("Failed to decode request body: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	sub := &models.Sub{
		ServiceName: req.ServiceName,
		Price: req.Price,
		UserID: req.UserID,
	}

	if err := h.subsService.UpdateSubService(sub, req.StartDate, req.EndDate); err != nil{
		utils.ErrorLogger.Printf("Failed to update sub: %v", err)
		http.Error(w, "failed to update subscription", http.StatusBadRequest)
		return
	}

	utils.InfoLogger.Printf("Sub updated successfully: %+v", sub)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

func (h* SubsHandler) DeleteSubHandler(w http.ResponseWriter, r *http.Request){
	utils.InfoLogger.Println("DeleteSubHandler called")
	id := r.URL.Query().Get("id")
	if id == ""{
		utils.WarningLogger.Println("Missing id parameter in request")
		http.Error(w, "missing id paramter", http.StatusBadRequest)
		return
	}

	if err := h.subsService.DeleteSubService(id); err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		utils.ErrorLogger.Printf("Failed to delete sub id=%s: %v", id, err)
		return
	}

	utils.InfoLogger.Printf("Sub Deleted successfully: id=%s", id)
	w.WriteHeader(http.StatusNoContent)
}
