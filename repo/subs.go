package repo

import (
	"online-subs-api/models"

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

