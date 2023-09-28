package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Survey struct {
	ID          int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OpenID      string    `json:"open_id" gorm:"column:open_id;type:int(11);not null;default:0;comment:'用户id'"`
	Content     string    `json:"content" gorm:"column:content;type:varchar(256);not null;comment:'内容'"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

func (obj *Survey) TableName() string {
	return "t_survey"
}

func (obj *Survey) Create(survey *Survey) (*Survey, error) {
	err := db.Get().Create(survey).Error
	return survey, err
}
