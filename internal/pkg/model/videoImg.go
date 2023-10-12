package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

// TVideoImg mapped from table <t_video_img>
type VideoImg struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"` // 自增主键
	ImgPath      string    `gorm:"column:img_path;comment:图片地址" json:"img_path"`                   // 图片地址
	CourtUUID    string    `gorm:"column:court_uuid;comment:场次uuid" json:"court_uuid"`             // 场次uuid
	CreateTime   time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`             // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`             // 更新时间
	RelativeTime int32     `gorm:"column:relative_time;comment:相对视频开始的时间戳" json:"relative_time"`   // 相对视频开始的时间戳
}

// TableName TVideoImg's table name
func (*VideoImg) TableName() string {
	return "t_video_img"
}

func (*VideoImg) Create(obj *VideoImg) (*VideoImg, error) {
	err := db.Get().Create(obj).Error
	return obj, err
}

func (obj *VideoImg) GetByCourtUuid(uuid string) ([]VideoImg, error) {
	results := make([]VideoImg, 0)
	err := db.Get().Table(obj.TableName()).Where("court_uuid = ?", uuid).Find(&results).Error
	return results, err
}
