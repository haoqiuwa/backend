package model

import (
	"time"
	"wxcloudrun-golang/internal/pkg/db"
)

type Order struct {
	ID          int32     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	OpenID      string    `gorm:"type:varchar(256);column:open_id;default:''" json:"open_id"`
	OrderType   int32     `gorm:"type:tinyint(2);column:order_type" json:"order_type"`
	Cost        float64   `gorm:"type:float;column:cost" json:"cost"`
	PaidCount   int32     `gorm:"column:paid_count" json:"paid_count"`
	BeforeCount int32     `gorm:"column:before_count" json:"before_count"`
	CreatedTime time.Time `gorm:"type:datetime;column:created_time;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time `gorm:"type:datetime;column:updated_time;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP" json:"updated_time"`
}

func (Order) TableName() string {
	return "t_order"
}

func (obj *Order) Create(order *Order) (*Order, error) {
	err := db.Get().Create(order).Error
	return order, err
}

// update
func (obj *Order) Update(order *Order) (*Order, error) {
	err := db.Get().Table(obj.TableName()).Where("id = ?", order.ID).Updates(order).Error
	return order, err
}

// GetByOpenID 根据openID获取用户信息
func (obj *Order) GetByOpenID(openID string) ([]*Order, error) {
	orders := make([]*Order, 0)
	err := db.Get().Where("open_id = ?", openID).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
