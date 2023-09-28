package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/resp"
)

// GetEvents 获取用户所属事件的视频
func (s *Service) GetEvents(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	courtID := c.Query("court")
	date := c.Query("date")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	dateInt, _ := strconv.Atoi(date)
	results, err := s.EventService.GetEvents(courtID, int32(dateInt))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(results, err))
}

// GetVideos 获取事件
func (s *Service) GetVideos(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	courtID := c.Query("court")
	date := c.Query("date")
	hour := c.Query("hour")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	hourInt, _ := strconv.Atoi(hour)
	dateInt, _ := strconv.Atoi(date)
	courtIDInt, _ := strconv.Atoi(courtID)
	event, err := s.EventService.GetVideos(int32(dateInt), int32(courtIDInt), int32(hourInt), openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(event, err))
}

// GetRecords 获取录像
func (s *Service) GetRecords(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	courtID := c.Query("court")
	date := c.Query("date")
	hour := c.Query("hour")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	hourInt, _ := strconv.Atoi(hour)
	dateInt, _ := strconv.Atoi(date)
	courtIDInt, _ := strconv.Atoi(courtID)
	data, err := s.EventService.GetRecord(int32(dateInt), int32(courtIDInt), int32(hourInt), openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// StoreVideo 存储视频
func (s *Service) StoreVideo(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	video := &model.Video{}
	err := json.Unmarshal(body, video)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	data, err := s.EventService.StoreVideo(video)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// GetMatchHighlights 获取比赛集锦
func (s *Service) GetMatchHighlights(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	courtID := c.Query("court")
	date := c.Query("date")
	hour := c.Query("hour")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	hourInt, _ := strconv.Atoi(hour)
	dateInt, _ := strconv.Atoi(date)
	courtIDInt, _ := strconv.Atoi(courtID)
	data, err := s.EventService.GetMatchHighlights(int32(dateInt), int32(courtIDInt), int32(hourInt), openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// GetMatchRecords 获取比赛录像
func (s *Service) GetMatchRecords(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	courtID := c.Query("court")
	date := c.Query("date")
	hour := c.Query("hour")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	hourInt, _ := strconv.Atoi(hour)
	dateInt, _ := strconv.Atoi(date)
	courtIDInt, _ := strconv.Atoi(courtID)
	data, err := s.EventService.GetMatchRecords(int32(dateInt), int32(courtIDInt), int32(hourInt), openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))

}

// GetAIContents 获取比赛录像
func (s *Service) GetAIContents(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	courtID := c.Query("court")
	date := c.Query("date")
	hour := c.Query("hour")
	if date == "" {
		date = time.Now().Format("20060102")
	}
	hourInt, _ := strconv.Atoi(hour)
	dateInt, _ := strconv.Atoi(date)
	courtIDInt, _ := strconv.Atoi(courtID)
	data, err := s.EventService.GetAIContent(int32(dateInt), int32(courtIDInt), int32(hourInt), openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))

}
