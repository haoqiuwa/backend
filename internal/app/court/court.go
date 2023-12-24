package court

import (
	"log"
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	courtDao *model.Court
}

func NewService() *Service {
	return &Service{
		courtDao: &model.Court{},
	}
}

type CourtWithDistance struct {
	model.Court
	Distance float64 `json:"distance"`
}

// GetCourts 获取所有场地，按距离倒序排列
func (s *Service) GetCourts(id int64) ([]model.Court, error) {
	results, err := s.courtDao.GetByVenueId(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return results, nil
}

// func sortByDistance(courts []CourtWithDistance) {
// 	for i := 0; i < len(courts); i++ {
// 		for j := i + 1; j < len(courts); j++ {
// 			if courts[i].Distance > courts[j].Distance {
// 				courts[i], courts[j] = courts[j], courts[i]
// 			}
// 		}
// 	}
// }

func (s *Service) GetAllCourts() ([]model.Court, error) {
	results, err := s.courtDao.Gets(&model.Court{})
	if err != nil {
		return nil, err
	}
	return results, nil
}

// func (s *Service) GetCountInfo(id int32, latitude, longitude string) (*CourtWithDistance, error) {
// 	court, err := s.courtDao.Get(&model.Court{ID: id})
// 	if err != nil {
// 		return nil, err
// 	}
// 	lat, _ := strconv.ParseFloat(latitude, 64)
// 	lng, _ := strconv.ParseFloat(longitude, 64)
// 	distance := location.GetDistance(lat, lng, court.Latitude, court.Longitude)
// 	result := &CourtWithDistance{
// 		Court:    *court,
// 		Distance: distance,
// 	}
// 	return result, nil
// }

// func (s *Service) JudgeLocation(courtID int64, latitude, longitude string) (bool, error) {
// 	court, err := s.courtDao.Get(&model.Court{ID: courtID})
// 	if err != nil {
// 		return false, err
// 	}
// 	lat, _ := strconv.ParseFloat(latitude, 64)
// 	lng, _ := strconv.ParseFloat(longitude, 64)
// 	distance := location.GetDistance(lat, lng, court.Latitude, court.Longitude)
// 	if distance <= 100 {
// 		return true, nil
// 	}
// 	return false, nil
// }

func (s *Service) GetCourtByID(id int32) (*model.Court, error) {
	court, err := s.courtDao.Get(&model.Court{ID: id})
	if err != nil {
		return nil, err
	}
	return court, nil
}
