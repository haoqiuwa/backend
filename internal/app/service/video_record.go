package service

import (
	"log"
	"strconv"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

type VideoRecordRes struct {
	UUID           string    `json:"uuid"`
	ID             int32     `json:"id"`
	FilePath       string    `json:"file_path"`
	Date           int32     `json:"date"`
	Time           int32     `json:"time"`
	VenueId        int32     `json:"venue_id"`
	CourtId        int32     `json:"court_id"`
	CourtCode      int32     `json:"court_code"`
	Hour           int32     `json:"hour"`
	CreatedTime    time.Time `json:"created_time"`
	UpdatedTime    time.Time `json:"updated_time"`
	StartTimestamp int64     `json:"start_timestamp"`
	EndTimestamp   int64     `json:"end_timestamp"`
	HoverImgPath   string    `json:"hover_img_path"`
	DownloadStatus bool      `json:"download_status"`
}

func (s *Service) GetVideoRecords(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
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
	r, err := s.VideoRecordService.GetVideoRecords(int32(venueId), int32(courtId), int32(date), int32(hour))
	log.Println("GetVideoRecords r", r)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	rd := s.FetchDownloadStatus(r, openID)
	c.JSON(200, resp.ToStruct(rd, err))
}

func (s *Service) FetchDownloadStatus(datas []*model.VideoRecord, openId string) []*VideoRecordRes {
	results := make([]*VideoRecordRes, 0)
	for _, v := range datas {
		rr := &VideoRecordRes{}
		rr.ID = v.ID
		rr.CourtCode = v.CourtCode
		rr.CourtId = v.CourtId
		rr.UUID = v.UUID
		rr.Date = v.Date
		rr.FilePath = v.FilePath
		rr.Time = v.Time
		rr.Hour = v.Hour
		rr.VenueId = v.VenueId
		rr.CreatedTime = v.CreatedTime
		rr.UpdatedTime = v.UpdatedTime
		rr.StartTimestamp = v.StartTimestamp
		rr.EndTimestamp = v.EndTimestamp
		rr.HoverImgPath = v.HoverImgPath
		d, err := s.DownloadRecordService.GetByOpenIdResourceIdAndresourceType(openId, v.ID, 20)
		if nil != err {
			log.Println("fetchDownloadStatus err", err)
		}
		if nil != d && d.ID > 0 {
			rr.DownloadStatus = true
		}
		results = append(results, rr)
	}
	return results
}
