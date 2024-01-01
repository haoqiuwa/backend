package videorecord

import (
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

func (s *Service) GetVideoRecords(venueId int32, courtId int32, date int32, hour int32) ([]*model.VideoRecord, error) {
	r, err := s.videorecordDao.GetVideoRecords(venueId, courtId, date, hour)
	if nil != err {
		return nil, err
	}
	return r, nil
}

func (s *Service) Create(v *model.VideoRecord) error {
	return s.videorecordDao.Create(v)
}

func (s *Service) GetById(id int32) (*model.VideoRecord, error) {
	return s.videorecordDao.Get(&model.VideoRecord{ID: id})
}
