package service

import (
	"io"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

func (s *Service) UnifiedOrder(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	ip := c.GetHeader("x-forwarded-for")
	body, _ := io.ReadAll(c.Request.Body)
	data, err := s.PayService.UnifiedOrder(openID, ip, string(body))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}
