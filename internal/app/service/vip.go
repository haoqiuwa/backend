package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/resp"
)

// GetVipCount	获取vip数量
func (s *Service) GetVipCount(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	data, err := s.VipService.GetRemainingCount(openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// GetVipOrders 获取vip订单
func (s *Service) GetVipOrders(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	data, err := s.VipService.GetOrdersByOpenID(openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// CreateVipOrder 创建vip订单
func (s *Service) CreateVipOrder(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	order := &model.Order{}
	err := json.Unmarshal(body, order)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	data, err := s.VipService.CreateOrder(openID, order.OrderType, order.Cost, order.PaidCount)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// UpdateVipCount 更新vip数量
func (s *Service) UpdateVipCount(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	order := &model.Vip{}
	err := json.Unmarshal(body, order)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	data, err := s.VipService.UpdateRemainingCount(openID, order.Count)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}
