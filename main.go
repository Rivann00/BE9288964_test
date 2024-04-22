package main

import "time"

type CatatanHarian struct {
	ID         int
	Tanggal    time.Time
	IsiCatatan string
}

type Repository interface {
	TambahCatatan(catatan CatatanHarian) error
	CariCatatan(tanggal time.Time) ([]CatatanHarian, error)
	HapusCatatan(id int) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) TambahCatatan(catatan CatatanHarian) error {
	return s.repo.TambahCatatan(catatan)
}

func (s *Service) CariCatatan(tanggal time.Time) ([]CatatanHarian, error) {
	return s.repo.CariCatatan(tanggal)
}

func (s *Service) HapusCatatan(id int) error {
	return s.repo.HapusCatatan(id)
}
