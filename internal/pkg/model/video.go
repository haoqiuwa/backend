package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Video struct {
	ID          int32     `gorm:"primary_key" json:"id"`
	FilePath    string    `gorm:"column:file_path" json:"file_path"`
	Date        int32     `gorm:"column:date" json:"date"`
	Time        string    `gorm:"column:time" json:"time"`
	Type        int32     `gorm:"column:type" json:"type"`
	Court       int32     `gorm:"column:court" json:"court"`
	Hour        int32     `gorm:"column:hour" json:"hour"`
	FileName    string    `gorm:"column:file_name" json:"file_name"`
	VideoName   string    `gorm:"column:video_name" json:"video_name"`
	FileType    int32     `gorm:"column:file_type" json:"file_type"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
}

// GORM table name for Video struct
func (obj *Video) TableName() string {
	return "t_video"
}
func (obj *Video) Create(video *Video) (*Video, error) {
	err := db.Get().Create(video).Error
	return video, err
}

func (obj *Video) Get(video *Video) (*Video, error) {
	result := new(Video)
	err := db.Get().Table(obj.TableName()).Where(video).First(result).Error
	return result, err
}

func (obj *Video) Gets(video *Video) ([]Video, error) {
	results := make([]Video, 0)
	err := db.Get().Table(obj.TableName()).Where(video).Find(&results).Error
	return results, err
}

func (obj *Video) Update(video *Video) (*Video, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", video.ID).Updates(video).Error
	return video, err
}

func (obj *Video) Delete(video *Video) error {
	return db.Get().Delete(video, "id = ?", video.ID).Error
}

func (obj *Video) GetDistinctHours(date int32) ([]int32, error) {
	results := make([]int32, 0)
	// get hours order by desc
	err := db.Get().Table(obj.TableName()).Where("date = ? and type in (1,2)",
		date).Order("hour desc").Pluck("distinct hour",
		&results).Error
	return results, err
}

func (obj *Video) GetVideos(date int32, courtID int32, hour int32, videoType int32) ([]*Video, error) {
	results := make([]*Video, 0)
	err := db.Get().Table(obj.TableName()).Where(
		"date = ? and court = ? and hour = ? and file_name like 'v%' and type = ?", date,
		courtID,
		hour, videoType).Order("file_name").Find(&results).Error
	return results, err
}

func (obj *Video) GetPictures(date int32, courtID int32, hour int32, videoType int32) ([]*Video, error) {
	results := make([]*Video, 0)
	err := db.Get().Table(obj.TableName()).Debug().Where(
		"date = ? and court = ? and hour = ? and file_name like 'p%'and type = ?", date,
		courtID,
		hour, videoType).Order("file_name").Find(&results).Error
	return results, err
}

func (obj *Video) GetMatchVideos(courtID int32, videoType int32) ([]*Video, error) {
	results := make([]*Video, 0)
	err := db.Get().Table(obj.TableName()).Where(
		"court = ? and file_name like 'v%' and type = ?",
		courtID,
		videoType).Order("created_time desc").Find(&results).Error
	return results, err
}

func (obj *Video) GetMatchPictures(courtID int32, videoType int32) ([]*Video, error) {
	results := make([]*Video, 0)
	err := db.Get().Table(obj.TableName()).Debug().Where(
		"court = ? and file_name like 'p%'and type = ?",
		courtID,
		videoType).Order("created_time desc").Find(&results).Error
	return results, err
}
