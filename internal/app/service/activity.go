package service

import (
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
