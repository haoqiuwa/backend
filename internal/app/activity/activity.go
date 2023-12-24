package activity

import (
	"log"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	ActivityDao     *model.Activity
	ActivityUserDao *model.ActivityUser
}

func NewService() *Service {
	return &Service{
		ActivityDao:     &model.Activity{},
		ActivityUserDao: &model.ActivityUser{},
	}
}

func (s *Service) CreateActivityUser(openId string, activityId int32, useId int32) (*model.ActivityUser, error) {
	return s.ActivityUserDao.Create(&model.ActivityUser{
		OpenID:     openId,
		ActivityID: activityId,
		UserID:     useId,
		CareteTime: time.Now(),
		UpdateTime: time.Now(),
	})
}

func (s *Service) FindActivityByIdAndOpenId(id int32, openId string) (*model.Activity, error) {
	log.Println("FindActivityByIdAndOpenId id openId", id, openId)
	a, err := s.ActivityDao.FindById(id)
	log.Println("FindActivityByIdAndOpenId activity", a)
	if nil != err {
		return nil, err
	}
	au, err := s.ActivityUserDao.FindByOpenIdAndActivityId(openId, id)
	log.Println("FindActivityByIdAndOpenId activity user ", au)
	if nil != au {
		return nil, err
	}
	return a, nil
}
