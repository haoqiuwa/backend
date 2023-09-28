package model

import (
	"gorm.io/gorm"
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Vip struct {
	ID          int32     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	OpenID      string    `gorm:"type:varchar(256);column:open_id;default:''" json:"open_id"`
	Count       int32     `gorm:"column:count" json:"count"`
	CreatedTime time.Time `gorm:"type:datetime;column:created_time;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time `gorm:"type:datetime;column:updated_time;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP" json:"updated_time"`
}

func (*Vip) TableName() string {
	return "t_vip"
}

// GetByOpenID 根据openID获取用户信息
func (obj *Vip) GetByOpenID(openID string) (*Vip, error) {
	var vip *Vip
	err := db.Get().Where("open_id = ?", openID).First(&vip).Error
	return vip, err
}

func (obj *Vip) Create(vip *Vip) (*Vip, error) {
	err := db.Get().Create(vip).Error
	return vip, err
}

// UpdateCountByOpenID 更新用户信息
func (obj *Vip) UpdateCountByOpenID(openID string, count int32) (*Vip, error) {
	vip, err := obj.GetByOpenID(openID)
	if err == gorm.ErrRecordNotFound {
		vip = &Vip{
			OpenID: openID,
			Count:  count,
		}
		_, err = obj.Create(vip)
		return vip, err
	}
	if err != nil {
		return nil, err
	}
	vip.Count = vip.Count + count
	err = db.Get().Save(vip).Error
	return vip, err
}
