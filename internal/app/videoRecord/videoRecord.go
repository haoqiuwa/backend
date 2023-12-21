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

func (s *Service) GetVideoRecords() ([]model.VideoRecord, error) {
	return s.videorecordDao.Gets(&model.VideoRecord{})
}

func (s *Service) Create(v *model.VideoRecord) (*model.VideoRecord, error) {
	return s.videorecordDao.Create(v)
}
