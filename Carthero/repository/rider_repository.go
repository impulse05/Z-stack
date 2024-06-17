package repository

import (
	"Carthero/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type RiderRepository interface {
	CreateRider(rider *model.Rider) error
	DeleteRider(id uint) error
	UpdateRiderState(id uint, assigned bool) error
	GetFreeRiders() ([]model.Rider, error)
}

type riderRepository struct {
	db *gorm.DB
}

func NewRiderRepository(db *gorm.DB) RiderRepository {
	return &riderRepository{db: db}
}

func (r *riderRepository) CreateRider(rider *model.Rider) error {
	return r.db.Create(rider).Error
}

func (r *riderRepository) DeleteRider(id uint) error {
	return r.db.Delete(&model.Rider{}, id).Error
}

func (r *riderRepository) UpdateRiderState(id uint, assigned bool) error {
	return r.db.Model(&model.Rider{}).Where("id = ?", id).Update("assigned", assigned).Error
}

func (r *riderRepository) GetFreeRiders() ([]model.Rider, error) {
	var riders []model.Rider
	err := r.db.Where("assigned = ?", false).Find(&riders).Error
	return riders, err
}
