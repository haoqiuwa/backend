package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Video struct {
	UUID         string    `gorm:"uuid" json:"uuid"`
	ID           int32     `gorm:"primary_key" json:"id"`
	FilePath     string    `gorm:"column:file_path" json:"file_path"`
	Date         int32     `gorm:"column:date" json:"date"`
	Time         int32     `gorm:"column:time" json:"time"`
	Type         int32     `gorm:"column:type" json:"type"`
	VenueId      int32     `gorm:"column:venue_id" json:"venue_id"`
	Court        int32     `gorm:"column:court" json:"court"`
	Hour         int32     `gorm:"column:hour" json:"hour"`
	FileName     string    `gorm:"column:file_name" json:"file_name"`
	VideoName    string    `gorm:"column:video_name" json:"video_name"`
	FileType     int32     `gorm:"column:file_type" json:"file_type"`
	CreatedTime  time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime  time.Time `gorm:"column:updated_time" json:"updated_time"`
	StartTime    int64     `gorm:"column:start_time" json:"start_time"`
	EndTime      int64     `gorm:"column:end_time" json:"end_time"`
	TeamAImgPath string    `gorm:"column:team_a_img_path" json:"team_a_img_path"`
	TeamBImgPath string    `gorm:"column:team_b_img_path" json:"team_b_img_path"`
	HoverImgPath string    `gorm:"column:hover_img_path" json:"hover_img_path"`
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

func (obj *Video) GetTimeRange(date int32) ([]int32, error) {
	results := make([]int32, 0)
	// get hours order by desc
	err := db.Get().Table(obj.TableName()).Where("date = ? and type =100",
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

func (obj *Video) GetVideoList(date int32, courtID int32, hour int32, venueId int32) ([]*Video, error) {
	results := make([]*Video, 0)
	err := db.Get().Table(obj.TableName()).Where(
		"date = ? and court = ? and hour = ? and venue_id = ? and type=100", date,
		courtID,
		hour, venueId).Order("id desc").Find(&results).Error
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

func (obj *Video) GetVideoByUUID(uuid string) (*Video, error) {
	result := new(Video)
	err := db.Get().Table(obj.TableName()).First(&result, "uuid=?", uuid).Error
	return result, err
}
func (obj *Video) GetVideoById(id int32) (*Video, error) {
	result := new(Video)
	err := db.Get().Table(obj.TableName()).First(&result, id).Error
	return result, err
}
