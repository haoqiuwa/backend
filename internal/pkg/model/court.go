package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Court struct {
	ID          int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CourtName   string    `json:"court_name" gorm:"column:court_name;type:varchar(255);default:'';comment:'场馆名称'"`
	VenueId     int32     `json:"venue_id" gorm:"column:venue_id;not null;comment:'场馆id'"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

// TableName get sql table name.获取数据库名字
func (obj *Court) TableName() string {
	return "t_court"
}

// Create 创建记录
func (obj *Court) Create(count *Court) (*Court, error) {
	err := db.Get().Create(count).Error
	return count, err
}

// Get 获取
func (obj *Court) Get(count *Court) (*Court, error) {
	result := new(Court)
	err := db.Get().Table(obj.TableName()).Where(count).First(result).Error
	return result, err
}

// Gets 获取批量结果
func (obj *Court) Gets(count *Court) ([]Court, error) {
	results := make([]Court, 0)
	err := db.Get().Table(obj.TableName()).Where(count).Find(&results).Error
	return results, err
}

// Update 更新
func (obj *Court) Update(count *Court) (*Court, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", count.ID).Updates(count).Error
	return count, err
}

// Delete 删除
func (obj *Court) Delete(count *Court) error {
	return db.Get().Delete(count, "id = ?", count.ID).Error
}

// GetsGetsWithLimit 获取批量结果
func (obj *Court) GetsWithLimit(count *Court, limit int32) ([]Court, error) {
	results := make([]Court, 0)
	err := db.Get().Table(obj.TableName()).Where(count).Limit(int(limit)).Find(&results).Error
	return results, err
}

func (obj *Court) GetByVenueId(id int64) ([]Court, error) {
	results := make([]Court, 0)
	err := db.Get().Table(obj.TableName()).Where("venue_id=?", id).Find(&results).Error
	return results, err
}
