package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

// TVideoClip mapped from table <t_video_clips>
type VideoClips struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`               // 自增主键
	FilePath     string    `gorm:"column:file_path;comment:视频地址" json:"file_path"`                               // 视频地址
	Team         string    `gorm:"column:team;comment:队" json:"team"`                                            // 队
	CourtUUID    string    `gorm:"column:court_uuid;comment:场次uuid" json:"court_uuid"`                           // 场次uuid
	HoverImgPath string    `gorm:"column:hover_img_path;comment:封面图地址" json:"hover_img_path"`                    // 封面图地址
	CreateTime   time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                           // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	VideoType    int32     `gorm:"column:video_type;default:1;comment:视频类型  1 集锦 2 ai视频" json:"video_type"`      // 视频类型  1 集锦 2 ai视频
	Time         int32     `gorm:"column:time;default:1;comment:视频时长  1 集锦 2 ai视频" json:"time"`                  // 视频时长
}

// TableName TVideoClip's table name
func (*VideoClips) TableName() string {
	return "t_video_clips"
}

func (*VideoClips) Create(obj *VideoClips) (*VideoClips, error) {
	err := db.Get().Create(obj).Error
	return obj, err
}

func (obj *VideoClips) GetByCourtUuid(uuid string) ([]VideoClips, error) {
	results := make([]VideoClips, 0)
	err := db.Get().Table(obj.TableName()).Where("court_uuid = ?", uuid).Find(&results).Error
	return results, err
}
func (obj *VideoClips) GetByCourtUuidAndVideoType(uuid string, videoType int32) ([]VideoClips, error) {
	results := make([]VideoClips, 0)
	err := db.Get().Table(obj.TableName()).Where("court_uuid = ? and video_type= ?", uuid, videoType).Find(&results).Error
	return results, err
}
