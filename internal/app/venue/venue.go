package venue

import (
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	venueDao *model.Venue
}

func NewService() *Service {
	return &Service{
		venueDao: &model.Venue{},
	}
}

func (s *Service) GetVenues() ([]model.Venue, error) {
	return s.venueDao.Gets(&model.Venue{})
}

func (s *Service) Create(v *model.Venue) (*model.Venue, error) {
	return s.venueDao.Create(v)
}

func (s *Service) GetVenueById(id int32) (*model.Venue, error) {
	return s.venueDao.Get(&model.Venue{ID: id})
}
