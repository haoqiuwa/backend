package download

import (
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	DownloadRecordDao *model.DownloadRecord
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Create(dr *model.DownloadRecord) (*model.DownloadRecord, error) {
	return s.DownloadRecordDao.Create(dr)
}

func (s *Service) GetById(id int32) (*model.DownloadRecord, error) {
	return s.DownloadRecordDao.GetById(id)
}

func (s *Service) GetByOpenIdAndType(openId string, resourceType int32) ([]model.DownloadRecord, error) {
	return s.DownloadRecordDao.GetByOpenIdAndType(openId, resourceType)
}

func (s *Service) GetByOpenIdPage(openId string, offset, page int32) ([]model.DownloadRecord, error) {
	return s.DownloadRecordDao.GetByOpenIdPage(openId, offset, page)
}
func (s *Service) GetByOpenIdResourceIdAndresourceType(openId string, resourceId, resourceType int32) (*model.DownloadRecord, error) {
	return s.DownloadRecordDao.GetByOpenIdResourceIdAndresourceType(openId, resourceId, resourceType)
}

// &model.DownloadRecord{
// 	OpenID:         "",
// 	ResourceUUID:   "",
// 	ResourceType:   1,
// 	CastDiamond:    2,
// 	FilePath:       "",
// 	HoverImgPath:   "",
// 	CurrentDiamond: 1,
// 	CourtId:        1,
// 	VenueId:        1,
// 	CourtName:      "",
// 	VenueName:      "",
// 	CreateTime:     time.Now(),
// 	UpdateTime:     time.Now(),
// }
