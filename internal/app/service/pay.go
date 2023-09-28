package service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"wxcloudrun-golang/internal/pkg/resp"
)

func (s *Service) UnifiedOrder(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	ip := c.GetHeader("x-forwarded-for")
	body, _ := ioutil.ReadAll(c.Request.Body)
	data, err := s.PayService.UnifiedOrder(openID, ip, string(body))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}
