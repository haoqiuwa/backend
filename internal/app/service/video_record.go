package service

import (
	"log"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetVideoRecords(c *gin.Context) {
	r, err := s.VideoRecordService.GetVideoRecords()
	log.Println("GetVideoRecords r", r)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(r, err))
}
