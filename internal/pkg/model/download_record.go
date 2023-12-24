package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

// DownloadRecord 视频下载记录
type DownloadRecord struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OpenID         string    `gorm:"column:open_id" json:"open_id"`
	ResourceId     int32     `gorm:"column:resource_id" json:"resource_id"`
	ResourceType   int32     `gorm:"column:resource_type" json:"resource_type"`
	ResourceUUID   string    `gorm:"column:resource_uuid" json:"resource_uuid"`
	CastDiamond    int32     `gorm:"column:cast_diamond" json:"cast_diamond"`
	FilePath       string    `gorm:"column:file_path" json:"file_path"`
	HoverImgPath   string    `gorm:"column:hover_img_path" json:"hover_img_path"`
	CurrentDiamond int32     `gorm:"column:current_diamond" json:"current_diamond"`
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`
	VenueId        int32     `gorm:"column:venue_id" json:"venue_id"`
	VenueName      string    `gorm:"column:venue_name" json:"venue_name"`
	CourtId        int32     `gorm:"column:court_id" json:"court_id"`
	CourtName      string    `gorm:"column:court_name" json:"court_name"`
}

// TableName TDownloadRecord's table name
func (*DownloadRecord) TableName() string {
	return "t_download_record"
}

func (obj *DownloadRecord) Create(model *DownloadRecord) (*DownloadRecord, error) {
	err := db.Get().Create(model).Error
	return model, err
}

func (obj *DownloadRecord) GetByResourceUUIDAndOpenId(resourceUUID, openId string) (*DownloadRecord, error) {
	r := &DownloadRecord{}
	err := db.Get().Table(obj.TableName()).Where("resource_uuid = ? and open_id = ?", resourceUUID, openId).Find(r).Error
	return r, err
}

func (obj *DownloadRecord) GetByOpenIdAndType(openId string, resourceType int32) ([]DownloadRecord, error) {
	results := make([]DownloadRecord, 0)
	err := db.Get().Table(obj.TableName()).Where("resource_type = ? and open_id = ?", resourceType, openId).Find(results).Error
	return results, err
}

func (obj *DownloadRecord) GetByOpenIdPage(openId string, offset, pageSize int32) ([]DownloadRecord, error) {
	results := make([]DownloadRecord, 0)
	err := db.Get().Table(obj.TableName()).Order("id desc").Offset(int(offset)).Limit(int(pageSize)).Where(" open_id = ?", openId).Find(&results).Error
	return results, err
}

func (obj *DownloadRecord) GetById(id int32) (*DownloadRecord, error) {
	r := &DownloadRecord{}
	err := db.Get().Table(obj.TableName()).Where("id = ?", id).Find(r).Error
	return r, err
}

func (obj *DownloadRecord) GetByOpenIdResourceIdAndresourceType(openid string, resourceId, resourceType int32) (*DownloadRecord, error) {
	r := &DownloadRecord{}
	err := db.Get().Table(obj.TableName()).Where("open_id = ? and resource_id= ? and resource_type=?", openid, resourceId, resourceType).Find(r).Error
	return r, err

}
