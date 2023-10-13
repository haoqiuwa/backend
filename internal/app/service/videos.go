package service

import (
	"encoding/json"
	"io"
	"log"
	"strconv"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/request"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
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
	body, _ := io.ReadAll(c.Request.Body)
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

// 处理硬件和算法端push过来的事件
func (s *Service) HandlePushEvent(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	log.Println("HandlePushEvent req:", string(body))
	req := &request.EventReq{}
	err := json.Unmarshal(body, req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	jsonb, err := json.Marshal(req.Data)
	if nil != err {
		c.JSON(400, err.Error())
		return
	}
	switch req.EventType {
	case 1:
		var eventData request.VideoEventReq
		err = toStruct(jsonb, &eventData)
		if nil != err {
			c.JSON(400, err.Error())
			return
		}
		log.Println("HandlePushEvent Received EventType1:", eventData)
		video := &model.Video{}
		video.Court = eventData.Court
		video.CreatedTime = time.Now()
		video.UpdatedTime = time.Now()
		video.Date = eventData.Date
		video.FileName = eventData.FileName
		video.FilePath = eventData.FilePath
		video.StartTime = eventData.StartTimestamp
		video.EndTime = eventData.EndTimestamp
		video.TeamAImgPath = eventData.TeamAImgPath
		video.TeamBImgPath = eventData.TeamBImgPath
		video.UUID = eventData.UUID
		video.HoverImgPath = eventData.HoverImgPath
		video.Time = eventData.Time
		jsonb, _ := json.Marshal(video)
		log.Println("json:", string(jsonb))
		_, err = s.EventService.StoreCourtVideo(video)
	case 2:
		var eventData request.VideoClipsEventReq
		err = toStruct(jsonb, &eventData)
		if nil != err {
			c.JSON(400, err.Error())
			return
		}
		videoClips := &model.VideoClips{}
		videoClips.CourtUUID = eventData.UUID
		videoClips.CreateTime = time.Now()
		videoClips.UpdateTime = time.Now()
		videoClips.HoverImgPath = eventData.HoverImgPath
		videoClips.FilePath = eventData.FilePath
		videoClips.VideoType = eventData.VideoType
		videoClips.Time = eventData.Time
		videoClips.Team = eventData.Team
		err = s.EventService.StoreVideoClips(videoClips)
		log.Println("HandlePushEvent Received EventType2 videoClips: ", videoClips)
	case 3:
		var eventData request.VideoImgEventReq
		err = toStruct(jsonb, &eventData)
		if nil != err {
			c.JSON(400, err.Error())
			return
		}
		vm := &model.VideoImg{}
		vm.CourtUUID = eventData.UUID
		vm.CreateTime = time.Now()
		vm.UpdateTime = time.Now()
		vm.ImgPath = eventData.FilePath
		log.Println("HandlePushEvent Received EventType3 VideoImg: ", vm)
		err = s.EventService.StoreVideoImg(vm)
	default:
		c.JSON(400, err.Error())
		return
	}
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct("ok", err))
}

func toStruct(jsonb []byte, s interface{}) error {
	err := json.Unmarshal(jsonb, s)
	if nil != err {
		log.Println("toStruct err", err)
	}
	return err
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

func (s *Service) GetAiVideos(c *gin.Context) {
	uuid := c.Param("uuid")
	data, err := s.EventService.GetAiVideos(uuid)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// 获取集锦视频
func (s *Service) GetHighlightsVideos(c *gin.Context) {
	uuid := c.Param("uuid")
	data, err := s.EventService.GetHighlightsVideos(uuid)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// 获取视频图册
func (s *Service) GetVideoImg(c *gin.Context) {
	uuid := c.Param("uuid")
	data, err := s.EventService.GetVideoImg(uuid)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}