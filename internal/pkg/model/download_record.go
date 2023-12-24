package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

// DownloadRecord 视频下载记录
type DownloadRecord struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OpenID         string    `gorm:"column:open_id" json:"open_id"`
	ResourceType   int32     `gorm:"column:resource_type" json:"resource_type"`
	ResourceUUID   string    `gorm:"column:resource_uuid" json:"resource_uuid"`
	CastDiamond    int32     `gorm:"column:cast_diamond" json:"cast_diamond"`
	FilePath       string    `gorm:"column:file_path" json:"file_path"`
	HoverImgPath   string    `gorm:"column:hover_img_path" json:"hover_img_path"`
	CurrentDiamond int32     `gorm:"column:current_diamond" json:"current_diamond"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName TDownloadRecord's table name
func (*DownloadRecord) TableName() string {
	return "t_download_record"
}

func (obj *DownloadRecord) Create(model *DownloadRecord) (*DownloadRecord, error) {
	err := db.Get().Create(model).Error
	return model, err
}
