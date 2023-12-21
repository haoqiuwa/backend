package videorecord

import (
	"log"
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	videorecordDao *model.VideoRecord
}

func NewService() *Service {
	return &Service{
		videorecordDao: &model.VideoRecord{},
	}
}

func (s *Service) GetVideoRecords(venueId int32, courtId int32, date int32, hour int32) ([]model.VideoRecord, error) {
	vr := &model.VideoRecord{}
	vr.Court = courtId
	vr.Date = date
	vr.VenueId = venueId
	vr.Hour = hour
	log.Println("GetVideoRecords", vr)
	return s.videorecordDao.Gets(vr)
}

func (s *Service) Create(v *model.VideoRecord) (*model.VideoRecord, error) {
	return s.videorecordDao.Create(v)
}
