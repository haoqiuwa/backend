package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wxcloudrun-golang/internal/pkg/resp"
)

// GetCounts 获取场地
func (s *Service) GetCounts(c *gin.Context) {
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	counts, err := s.CourtService.GetCourts(latitude, longitude)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(counts, err))
}

func (s *Service) GetCountInfo(c *gin.Context) {
	countID := c.Param("id")
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	countIDInt, _ := strconv.Atoi(countID)
	countInfo, err := s.CourtService.GetCountInfo(int32(countIDInt), latitude, longitude)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(countInfo, err))
}

// GetRecommendVideos 获取推荐视频
func (s *Service) GetRecommendVideos(c *gin.Context) {
	videos, err := s.RecommendService.GetRecommend()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(videos, err))
}

// JudgeLocation 判断用户是否在场地内
func (s *Service) JudgeLocation(c *gin.Context) {
	countID := c.Param("id")
	countIDInt, _ := strconv.Atoi(countID)
	latitude := c.Query("latitude")
	longitude := c.Query("longitude")
	result, err := s.CourtService.JudgeLocation(int32(countIDInt), latitude, longitude)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(result, err))
}
