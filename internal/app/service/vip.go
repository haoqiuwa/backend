package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/request"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

// GetVipCount	获取vip数量
func (s *Service) GetVipCount(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	log.Println("GetVipCount openID:", openID)
	if openID == "" {
		c.JSON(http.StatusBadRequest, "请先登录")
		return
	}
	data, err := s.VipService.GetRemainingCount(openID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp.ToStruct(data, err))
}

// GetVipCount	获取vip数量
func (s *Service) GetVipInfo(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	log.Println("GetVipCount openID:", openID)
	if openID == "" {
		c.JSON(http.StatusBadRequest, "请先登录")
		return
	}
	data, err := s.VipService.GetByOpenID(openID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp.ToStruct(data, err))
}

// GetVipOrders 获取vip订单
func (s *Service) GetVipOrders(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(http.StatusBadRequest, "请先登录")
		return
	}
	data, err := s.VipService.GetOrdersByOpenID(openID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp.ToStruct(data, err))
}

// CreateVipOrder 创建vip订单
func (s *Service) CreateVipOrder(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(http.StatusBadRequest, "请先登录")
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	order := &model.Order{}
	err := json.Unmarshal(body, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	data, err := s.VipService.CreateOrder(openID, order.OrderType, order.Cost, order.PaidCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp.ToStruct(data, err))
}

// UpdateVipCount 更新vip数量
func (s *Service) UpdateVipCount(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(http.StatusBadRequest, "请先登录")
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	order := &request.UpdateVipCountCnt{}
	err := json.Unmarshal(body, order)
	if order.Count < 0 && order.FilePath != "" {
		data, _ := s.CollectService.GetUserDownloadStatus(openID, order.FilePath)
		if data {
			c.JSON(http.StatusOK, resp.ToStruct(data, err))
		}
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	data, err := s.VipService.UpdateRemainingCount(openID, order.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp.ToStruct(data, err))
}
