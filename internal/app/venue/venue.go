package venue

import (
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	venueDao *model.Venue
}

func (s *Service) GetVenues() ([]model.Venue, error) {
	return s.venueDao.Gets(&model.Venue{})
}
