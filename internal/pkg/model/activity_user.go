package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

// ActivityUser 参与活动用户表
type ActivityUser struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	OpenID     string    `gorm:"column:open_id;comment:用户openid" json:"open_id"`               // 用户openid
	UserID     int32     `gorm:"column:user_id;comment:用户id" json:"user_id"`                   // 用户id
	ActivityID int64     `gorm:"column:activity_id;comment:活动id" json:"activity_id"`           // 活动id
	CareteTime time.Time `gorm:"column:carete_time;comment:创建时间" json:"carete_time"`           // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`           // 更新时间
}

// TableName ActivityUser's table name
func (*ActivityUser) TableName() string {
	return "t_activity_user"
}

func (obj *ActivityUser) FindByOpenIdAndActivityId(openId string, activityId int32) (*ActivityUser, error) {
	ac := &ActivityUser{}
	err := db.Get().Table(obj.TableName()).Where("open_id = ? and activity_id = ?", openId, activityId).Find(ac).Error
	return ac, err
}
