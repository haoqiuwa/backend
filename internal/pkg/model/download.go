package model

import (
	"fmt"
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Download struct {
	ID          int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	OpenID      string    `json:"open_id" gorm:"column:open_id;type:int(11);not null;default:0;comment:'用户id'"`
	FileID      string    `json:"file_id" gorm:"column:file_id;type:varchar(256);not null;comment:'视频文件id'"`
	CreatedTime Time      `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

// 自定义 Time 类型
type Time time.Time

// 实现 MarshalJSON 方法
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format("2006-01-02 15:04:05"))), nil
}
func (obj *Download) TableName() string {
	return "t_download"
}

func (obj *Download) Create(collect *Download) (*Download, error) {
	err := db.Get().Create(collect).Error
	return collect, err
}

func (obj *Download) Get(collect *Download) (*Download, error) {
	result := new(Download)
	err := db.Get().Table(obj.TableName()).Where(collect).First(result).Error
	return result, err
}

func (obj *Download) Gets(collect *Download) ([]Download, error) {
	results := make([]Download, 0)
	err := db.Get().Table(obj.TableName()).Where(collect).Find(&results).Error
	return results, err
}

func (obj *Download) Update(collect *Download) (*Download, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", collect.ID).Updates(collect).Error
	return collect, err
}

func (obj *Download) Delete(collect *Download) error {
	return db.Get().Delete(collect, "id = ?", collect.ID).Error
}
