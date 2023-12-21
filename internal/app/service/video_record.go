package service

import (
	"log"
	"strconv"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetVideoRecords(c *gin.Context) {
	dateStr := c.Query("date")
	date, err := strconv.Atoi(dateStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	hourStr := c.Query("hour")
	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	courtIdStr := c.Query("courtId")
	courtId, err := strconv.Atoi(courtIdStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	venueIdStr := c.Query("venueId")
	venueId, err := strconv.Atoi(venueIdStr)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	r, err := s.VideoRecordService.GetVideoRecords(int32(date), int32(courtId), int32(hour), int32(venueId))
	log.Println("GetVideoRecords r", r)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(r, err))
}
