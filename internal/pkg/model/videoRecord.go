package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type VideoRecord struct {
	UUID         string    `gorm:"uuid" json:"uuid"`
	ID           int32     `gorm:"primary_key" json:"id"`
	FilePath     string    `gorm:"column:file_path" json:"file_path"`
	Date         int32     `gorm:"column:date" json:"date"`
	Time         int32     `gorm:"column:time" json:"time"`
	VenueId      int32     `gorm:"column:venue_id" json:"venue_id"`
	Court        int32     `gorm:"column:court" json:"court"`
	Hour         int32     `gorm:"column:hour" json:"hour"`
	CreatedTime  time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime  time.Time `gorm:"column:updated_time" json:"updated_time"`
	StartTime    int64     `gorm:"column:start_time" json:"start_time"`
	EndTime      int64     `gorm:"column:end_time" json:"end_time"`
	HoverImgPath string    `gorm:"column:hover_img_path" json:"hover_img_path"`
}

// GORM table name for Video struct
func (obj *VideoRecord) TableName() string {
	return "t_video_record"
}

func (obj *VideoRecord) Create(video *VideoRecord) (*VideoRecord, error) {
	err := db.Get().Create(video).Error
	return video, err
}

func (obj *VideoRecord) Get(video *VideoRecord) (*VideoRecord, error) {
	result := new(VideoRecord)
	err := db.Get().Table(obj.TableName()).Where(video).First(result).Error
	return result, err
}

func (obj *VideoRecord) Gets(video *VideoRecord) ([]VideoRecord, error) {
	results := make([]VideoRecord, 0)
	err := db.Get().Table(obj.TableName()).Where(video).Find(&results).Error
	return results, err
}

func (obj *VideoRecord) Update(video *VideoRecord) (*VideoRecord, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", video.ID).Updates(video).Error
	return video, err
}
