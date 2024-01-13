package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Venue struct {
	ID          int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	VenueName   string    `json:"venue_name" gorm:"column:venue_name;type:varchar(255);default:'';comment:'场馆名称'"`
	VenueAddr   string    `json:"venue_addr" gorm:"column:venue_addr;type:text;default:'';comment:'场馆地址'"`
	VenueConf   string    `json:"venue_conf" gorm:"column:venue_conf;type:text;default:'';comment:'场馆配置json'"`
	VenueDesc   string    `json:"venue_desc" gorm:"column:venue_desc;type:text;default:'';comment:'场馆描述'"`
	VenueLogo   string    `json:"venue_logo" gorm:"column:venue_logo;type:varchar(255);default:'';comment:'场馆logo'"`
	Latitude    string    `json:"latitude" gorm:"column:latitude;type:varchar(255);default:'';comment:'场馆经纬度'"`
	Longitude   string    `json:"longitude" gorm:"column:longitude;type:varchar(255);default:'';comment:'场馆经纬度'"`
	VenuePhone  string    `json:"venue_phone" gorm:"column:venue_phone;type:varchar(255);default:'';comment:'场馆联系电话'"`
	CreateUser  string    `json:"create_user" gorm:"column:create_user;type:varchar(255);default:'';comment:'创建人'"`
	UpdateUser  string    `json:"update_user" gorm:"column:update_user;type:varchar(255);default:'';comment:'更新人'"`
	VenueStatus bool      `json:"venue_status" gorm:"column:venue_status;type:tinyint(1);default:1;comment:'场馆状态"`
	IsDel       bool      `json:"is_del" gorm:"column:is_del;type:tinyint(1);default:1;comment:'是否删除'"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
	ShortName   string    `json:"short_name" gorm:"column:short_name;type:varchar(255);default:'';comment:'简称'"`
	ServiceTime string    `json:"service_time" gorm:"column:service_time;type:varchar(255);default:'';comment:'营业时间'"`
}

// TableName get sql table name.获取数据库名字
func (obj *Venue) TableName() string {
	return "t_venue"
}

// Create 创建记录
func (obj *Venue) Create(v *Venue) (*Venue, error) {
	err := db.Get().Create(v).Error
	return v, err
}

// Get 获取
func (obj *Venue) Get(v *Venue) (*Venue, error) {
	result := new(Venue)
	err := db.Get().Table(obj.TableName()).Where(v).First(result).Error
	return result, err
}

// Get 获取
func (obj *Venue) GetList() ([]Venue, error) {
	results := make([]Venue, 0)
	err := db.Get().Table(obj.TableName()).Where("id > 0").Find(&results).Error
	return results, err
}

// Gets 获取批量结果
func (obj *Venue) Gets(v *Venue) ([]Venue, error) {
	results := make([]Venue, 0)
	err := db.Get().Table(obj.TableName()).Where(v).Find(&results).Error
	return results, err
}

// Update 更新
func (obj *Venue) Update(v *Venue) (*Venue, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", v.ID).Updates(v).Error
	return v, err
}

// Delete 删除
func (obj *Venue) Delete(v *Venue) error {
	return db.Get().Delete(v, "id = ?", v.ID).Error
}

// GetsGetsWithLimit 获取批量结果
func (obj *Venue) GetsWithLimit(count *Venue, limit int32) ([]Venue, error) {
	results := make([]Venue, 0)
	err := db.Get().Table(obj.TableName()).Where(count).Limit(int(limit)).Find(&results).Error
	return results, err
}
