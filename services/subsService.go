package services

import (
	"errors"
	"fmt"
	"online-subs-api/models"
	"online-subs-api/repo"
	"online-subs-api/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SubsService struct{
	subsRepo *repo.SubsRepo
}

func NewSubsService(subsRepo *repo.SubsRepo) *SubsService{
	return &SubsService{subsRepo: subsRepo}
}

func validateUUID(id string) bool{
	re := regexp.MustCompile(`^[a-fA-F0-9\\-]{36}$`)
	return re.MatchString(id)
}

//parse json date into time.Time date ("07-2025" -> "2025-07-01")
func validDate(date string) (time.Time, error){
	parts := strings.Split(date, "-")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("invalid date format in JSON, expected MM-YYYY")
	}

	month, err := strconv.Atoi(parts[0])
	if err != nil || month < 1 || month > 12 {
		return time.Time{}, fmt.Errorf("invalid month, should be between [1-12]")
	}

	year, err := strconv.Atoi(parts[1])
	if err != nil || year < 1 {
		return time.Time{}, fmt.Errorf("invalid year")
	}

	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC), nil
}

func (s *SubsService) CreateService(sub *models.Sub, startDateStr, endDateStr string) error{
	if !validateUUID(sub.UserID){
		utils.ErrorLogger.Println("Invalid user_id format:", sub.UserID)
		return errors.New("invalid user_id format")
	}

	if sub.Price <= 0{
		utils.ErrorLogger.Println("Invalid price provided:", sub.Price)
		return errors.New("price must be a postive integer")
	}

	startDate, err := validDate(startDateStr)
	if err != nil {
		utils.ErrorLogger.Println("Invalid start date:", startDateStr, "error:", err)
		return err
	}
	sub.StartDate = startDate

	if endDateStr != ""{
		endDate, err := validDate(endDateStr)
		if err != nil {
			utils.ErrorLogger.Println("Invalid end date:", endDateStr, "error:", err)
			return err
		}
		sub.EndDate = endDate
	}

	id, err := utils.NewUUID()
	if err != nil {
		utils.ErrorLogger.Println("Failed to generate UUID:", err)
		return err
	}
	sub.ID = id
	return s.subsRepo.CreateSubRepo(sub)
}

func (s *SubsService) GetServiceByID(id string) (*models.Sub, error){
	if !validateUUID(id) {
		utils.ErrorLogger.Println("Invalid ID format:", id)
		return nil, errors.New("invalid id format")
	}
	return s.subsRepo.GetSubRepoById(id)
}

func (s *SubsService) ListAllSubsService(userID, serviceName string) ([]models.Sub, error){
	if userID != "" && !validateUUID(userID) {
		utils.ErrorLogger.Println("Invalid user_id format:", userID)
		return nil, errors.New("invalid user_id format")
	}
	return s.subsRepo.ListAllSubsRepo(userID, serviceName)
}

func (s *SubsService) UpdateSubService(sub *models.Sub, startDateStr, endDateStr string) error{
	if !validateUUID(sub.UserID){
		utils.ErrorLogger.Println("Invalid user_id format:", sub.UserID)
		return errors.New("invalid user_id format")
	}

	if sub.Price <= 0{
		utils.ErrorLogger.Println("Invalid price provided:", sub.Price)
		return errors.New("price must be a postive integer")
	}

	startDate, err := validDate(startDateStr)
	if err != nil {
		utils.ErrorLogger.Println("Invalid start date:", startDateStr, "error:", err)
		return err
	}
	sub.StartDate = startDate

	if endDateStr != ""{
		endDate, err := validDate(endDateStr)
		if err != nil {
			utils.ErrorLogger.Println("Invalid end date:", endDateStr, "error:", err)
			return err
		}
		sub.EndDate = endDate
	}

	id, err := utils.NewUUID()
	if err != nil {
		utils.ErrorLogger.Println("Failed to generate UUID:", err)
		return err
	}
	sub.ID = id
	return s.subsRepo.UpdateSubRepo(sub)
}

func (s *SubsService) DeleteSubService(id string) error{
	if !validateUUID(id) {
		utils.ErrorLogger.Println("Invalid ID format:", id)
		return errors.New("invalid id format")
	}
	return s.subsRepo.DeleteSubRepo(id)
}


func (s *SubsService) GetTotalCostService(startStr, endStr, userID, serviceName string) (int, error) {
	start, err := validDate(startStr)
	if err != nil {
		return 0, err
	}
	end, err := validDate(endStr)
	if err != nil {
		return 0, err
	}

	if userID != "" && !validateUUID(userID) {
		return 0, errors.New("invalid user_id format")
	}

	return s.subsRepo.GetTotalCostRepo(start, end, userID, serviceName)
}
