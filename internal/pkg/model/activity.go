package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

// Activity 活动表
type Activity struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:主健" json:"id"`       // 主健
	ActivityType    int32     `gorm:"column:activity_type;comment:类型" json:"activity_type"`               // 类型
	ActivityContext string    `gorm:"column:activity_context;comment:内容json" json:"activity_context"`     // 内容json
	ActivityConfig  string    `gorm:"column:activity_config;comment:配置json" json:"activity_config"`       // 配置json
	ActivityStatus  bool      `gorm:"column:activity_status;default:1;comment:状态" json:"activity_status"` // 状态
	StartTime       time.Time `gorm:"column:start_time;comment:开始时间" json:"start_time"`                   // 开始时间
	EndTime         time.Time `gorm:"column:end_time;comment:结束时间" json:"end_time"`                       // 结束时间
	CreateTime      time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                 // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                 // 更新时间
}

// TableName Activity's table name
func (*Activity) TableName() string {
	return "t_activity"
}

func (obj *Activity) FindById(id int32) (*Activity, error) {
	r := &Activity{}
	err := db.Get().Table(obj.TableName()).Find(r, id).Error
	return r, err
}
