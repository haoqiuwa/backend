package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type VideoRecord struct {
	UUID           string    `gorm:"uuid" json:"uuid"`
	ID             int32     `gorm:"primary_key" json:"id"`
	FilePath       string    `gorm:"column:file_path" json:"file_path"`
	Date           int32     `gorm:"column:date" json:"date"`
	Time           int32     `gorm:"column:time" json:"time"`
	VenueId        int32     `gorm:"column:venue_id" json:"venue_id"`
	CourtId        int32     `gorm:"column:court_id" json:"court_id"`
	CourtCode      int32     `gorm:"column:court_code" json:"court_code"`
	Hour           int32     `gorm:"column:hour" json:"hour"`
	CreatedTime    time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime    time.Time `gorm:"column:updated_time" json:"updated_time"`
	StartTimestamp int64     `gorm:"column:start_timestamp" json:"start_timestamp"`
	EndTimestamp   int64     `gorm:"column:end_timestamp" json:"end_timestamp"`
	HoverImgPath   string    `gorm:"column:hover_img_path" json:"hover_img_path"`
}

// GORM table name for Video struct
func (obj *VideoRecord) TableName() string {
	return "t_video_record"
}

func (obj *VideoRecord) Create(vr *VideoRecord) error {
	err := db.Get().Create(vr).Error
	return err
}

func (obj *VideoRecord) Get(vr *VideoRecord) (*VideoRecord, error) {
	result := new(VideoRecord)
	err := db.Get().Table(obj.TableName()).Where(vr).First(result).Error
	return result, err
}

func (obj *VideoRecord) Gets(vr *VideoRecord) ([]VideoRecord, error) {
	results := make([]VideoRecord, 0)
	err := db.Get().Table(obj.TableName()).Where(vr).Find(&results).Error
	return results, err
}

func (obj *VideoRecord) Update(vr *VideoRecord) (*VideoRecord, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", vr.ID).Updates(vr).Error
	return vr, err
}

func (obj *VideoRecord) GetVideoRecords(venueId int32, courtId int32, date int32, hour int32) ([]*VideoRecord, error) {
	results := make([]*VideoRecord, 0)
	err := db.Get().Table(obj.TableName()).Where("venue_id = ?  and court_id = ? and date = ?  and hour = ?  ", venueId, courtId, date, hour).Order("id desc").Find(&results).Error
	return results, err
}
