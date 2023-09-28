package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"

	"gorm.io/gorm/clause"
)

type User struct {
	ID          int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	OpenID      string    `json:"open_id" gorm:"column:open_id;type:varchar(255);not null;unique_index"`
	Phone       string    `json:"phone" gorm:"column:phone;type:varchar(255);not null;unique_index"`
	Court       int32     `json:"court" gorm:"column:court;type:int(11);not null;default:0;comment:'球场'"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'创建时间'"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:'更新时间'"`
}

// TableName get sql table name.获取数据库名字
func (obj *User) TableName() string {
	return "t_user"
}

// Create 创建记录
func (obj *User) Create(user *User) (*User, error) {
	// create, if record exists, update
	err := db.Get().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(user).Error
	return user, err
}

// Get 获取
func (obj *User) Get(user *User) (*User, error) {
	result := new(User)
	err := db.Get().Table(obj.TableName()).Where(user).First(result).Error
	return result, err
}

// Gets 获取批量结果
func (obj *User) Gets(user *User) ([]User, error) {
	results := make([]User, 0)
	err := db.Get().Table(obj.TableName()).Where(user).Find(&results).Error
	return results, err
}

// Update 更新
func (obj *User) Update(user *User) (*User, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", user.ID).Updates(user).Error
	return user, err
}

// Delete 删除
func (obj *User) Delete(user *User) error {
	return db.Get().Delete(user, "id = ?", user.ID).Error
}
