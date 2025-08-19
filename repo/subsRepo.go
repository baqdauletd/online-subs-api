package repo

import (
	"online-subs-api/models"
	"time"

	"gorm.io/gorm"
)

type SubsRepo struct{
	db *gorm.DB
}

func NewSubsRepo(db *gorm.DB) *SubsRepo{
	return &SubsRepo{
		db: db,
	}
}

func (r *SubsRepo) CreateSubRepo (subs *models.Sub) error{
	return	r.db.Create(subs).Error
}

func (r *SubsRepo) GetSubRepoById(id string) (*models.Sub, error){
	var sub models.Sub
	if err := r.db.First(&sub, "id=?", id).Error; err != nil{
		return nil, err
	}
	return &sub, nil
}

func (r *SubsRepo) ListAllSubsRepo(userID, serviceName string) ([]models.Sub, error){
	var subs []models.Sub
	query := r.db

	if userID != ""{
		query = query.Where("user_id=?", userID)
	}
	if serviceName != ""{
		query = query.Where("service_name=?", serviceName)
	}

	if err := query.Find(&subs).Error; err != nil{
		return nil, err
	}

	return subs, nil
}

func (r *SubsRepo) UpdateSubRepo(sub *models.Sub) error{
	return r.db.Save(sub).Error
}

func (r *SubsRepo) DeleteSubRepo(id string) error{
	return r.db.Delete(&models.Sub{}, "id=?", id).Error
}

func (r *SubsRepo) GetTotalCostRepo(startDate, endDate time.Time, userID, serviceName string) (int, error) {
	var total int64
	query := r.db.Model(&models.Sub{})

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}

	query = query.Where(`
		start_date <= ? 
		AND (end_date IS NULL OR end_date >= ?)`,
		endDate, startDate,
	)

	if err := query.Select("SUM(price)").Scan(&total).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}

