package services

import (
	"Carthero/model"
	"Carthero/repository"
)

type RiderService interface {
	CreateRider(rider *model.Rider) error
	DeleteRider(id uint) error
	UpdateRiderState(id uint, assigned bool) error
	GetFreeRiders() ([]model.Rider, error)
}

type riderService struct {
	riderRepo repository.RiderRepository
}

func NewRiderService(riderRepo repository.RiderRepository) RiderService {
	return &riderService{riderRepo: riderRepo}
}

func (s *riderService) CreateRider(rider *model.Rider) error {
	return s.riderRepo.CreateRider(rider)
}

func (s *riderService) DeleteRider(id uint) error {
	return s.riderRepo.DeleteRider(id)
}

func (s *riderService) UpdateRiderState(id uint, assigned bool) error {
	return s.riderRepo.UpdateRiderState(id, assigned)
}

func (s *riderService) GetFreeRiders() ([]model.Rider, error) {
	return s.riderRepo.GetFreeRiders()
}
