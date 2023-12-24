package service

import (
	"encoding/json"
	"io"
	"wxcloudrun-golang/internal/pkg/request"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

const id int32 = 1

// 获取活动 如果用户参加过活动了则不返回
func (s *Service) GetActivity(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	a, err := s.ActivityService.FindActivityByIdAndOpenId(id, openID)
	if err != nil {
		c.JSON(200, "暂无活动")
		return
	}
	c.JSON(200, resp.ToStruct(a, err))
}

// 获取活动 如果用户参加过活动了则不返回
func (s *Service) UseActivity(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	v, err := s.VipService.GetByOpenID(openID)
	if nil != err {
		c.JSON(400, err.Error())
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	useActivity := &request.UseActivity{}
	err = json.Unmarshal(body, useActivity)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	activity, err := s.ActivityService.FindActivityById(useActivity.ActivityId)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}
	ac, err := s.ActivityService.CreateActivityUser(openID, activity.ID, v.ID)
	if nil != err {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(ac, err))
}
