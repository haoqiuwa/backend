package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type UserEvent struct {
	ID          int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OpenID      string    `json:"open_id" gorm:"column:open_id;type:int(11);not null;default:0;comment:'用户id'"`
	FileID      string    `json:"file_id" gorm:"column:file_id;type:varchar(256);not null;comment:'视频文件id'"`
	EventType   int32     `json:"event_type" gorm:"column:event_type;type:int(11);not null;default:0;comment:'事件类型'"`
	FromPage    string    `json:"from_page" gorm:"column:from_page;type:varchar(256);not null;comment:'前置页'"`
	VideoType   int32     `json:"video_type" gorm:"column:video_type;type:int(11);not null;default:0;comment:'视频类型'"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

func (obj *UserEvent) TableName() string {
	return "t_user_event"
}

func (obj *UserEvent) Create(userEvent *UserEvent) (*UserEvent, error) {
	err := db.Get().Create(userEvent).Error
	return userEvent, err
}

func (obj *UserEvent) Get(userEvent *UserEvent) (*UserEvent, error) {
	result := new(UserEvent)
	err := db.Get().Table(obj.TableName()).Where(userEvent).First(result).Error
	return result, err
}

func (obj *UserEvent) Gets(userEvent *UserEvent) ([]UserEvent, error) {
	results := make([]UserEvent, 0)
	err := db.Get().Table(obj.TableName()).Where(userEvent).Find(&results).Error
	return results, err
}
