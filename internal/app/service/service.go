package service

import (
	"wxcloudrun-golang/internal/app/collect"
	"wxcloudrun-golang/internal/app/court"
	"wxcloudrun-golang/internal/app/event"
	"wxcloudrun-golang/internal/app/pay"
	"wxcloudrun-golang/internal/app/recommend"
	"wxcloudrun-golang/internal/app/user"
	"wxcloudrun-golang/internal/app/vip"
)

type Service struct {
	UserService      *user.Service
	CourtService     *court.Service
	EventService     *event.Service
	CollectService   *collect.Service
	RecommendService *recommend.Service
	VipService       *vip.Service
	PayService       *pay.Service
}

func NewService() *Service {
	return &Service{
		UserService:      user.NewService(),
		CourtService:     court.NewService(),
		EventService:     event.NewService(),
		CollectService:   collect.NewService(),
		RecommendService: recommend.NewService(),
		VipService:       vip.NewService(),
		PayService:       pay.NewService(),
	}
}
