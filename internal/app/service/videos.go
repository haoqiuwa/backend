package service

import (
	"encoding/json"
	"io"
	"log"
	"sort"
	"strconv"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/request"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

type VideoClipsRes struct {
	ID             int32     `json:"id"`             // 自增主键
	FilePath       string    `json:"file_path"`      // 视频地址
	Team           string    `json:"team"`           // 队
	CourtUUID      string    `json:"court_uuid"`     // 场次uuid
	HoverImgPath   string    `json:"hover_img_path"` // 封面图地址
	CreateTime     time.Time `json:"create_time"`    // 创建时间
	UpdateTime     time.Time `json:"update_time"`    // 更新时间
	VideoType      int32     `json:"video_type"`     // 视频类型  1 集锦 2 ai视频
	Time           int32     `json:"time"`           // 视频时长
	TimeRange      string    `json:"time_range"`
	DownloadStatus bool      `json:"download_status"`
}

type VideoRes struct {
	UUID           string    `json:"uuid"`
	ID             int32     `json:"id"`
	FilePath       string    `json:"file_path"`
	Date           int32     `json:"date"`
	Time           int32     `json:"time"`
	Type           int32     `json:"type"`
	VenueId        int32     `json:"venue_id"`
	Court          int32     `json:"court"`
	Hour           int32     `json:"hour"`
	FileName       string    `json:"file_name"`
	VideoName      string    `json:"video_name"`
	FileType       int32     `json:"file_type"`
	CreatedTime    time.Time `json:"created_time"`
	UpdatedTime    time.Time `json:"updated_time"`
	StartTime      int64     `json:"start_time"`
	EndTime        int64     `json:"end_time"`
	TeamAImgPath   string    `json:"team_a_img_path"`
	TeamBImgPath   string    `json:"team_b_img_path"`
	HoverImgPath   string    `json:"hover_img_path"`
	DownloadStatus bool      `json:"download_status"`
}

type TimeRangeRes struct {
	VenueId        int32 `json:"venue_id"`
	CourtId        int32 `json:"court_id"`
	CourtCode      int32 `json:"court_code"`
	Date           int32 `json:"date"`
	Hour           int32 `json:"hour"`
	VideoCnt       int32 `json:"video_cnt"`
	VideoRecordCnt int32 `json:"video_record_cnt"`
	VideoClipsCnt  int32 `json:"video_clips_cnt"`
}

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

// 添加日志出发发布
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
		video.VenueId = eventData.VenueId
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
		video.Type = 100
		video.Time = eventData.Time
		video.Hour = eventData.Hour
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
		videoClips.TimeRange = eventData.TimeRange
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
		vm.ImgType = eventData.ImgType
		log.Println("HandlePushEvent Received EventType3 VideoImg: ", vm)
		err = s.EventService.StoreVideoImg(vm)
	case 4:
		var eventData request.VideoRecordEventReq
		err = toStruct(jsonb, &eventData)
		if nil != err {
			c.JSON(400, err.Error())
			return
		}
		log.Println("HandlePushEvent Received EventType4:", eventData)
		ct, err := s.CourtService.GetByVenueIdAndCode(eventData.VenueId, eventData.Court)
		if nil != err {
			c.JSON(400, err.Error())
			return
		}
		video := &model.VideoRecord{}
		video.CourtCode = eventData.Court
		video.CourtId = ct.ID
		video.VenueId = eventData.VenueId
		video.CreatedTime = time.Now()
		video.UpdatedTime = time.Now()
		video.Date = eventData.Date
		video.FilePath = eventData.FilePath
		video.StartTimestamp = eventData.StartTimestamp
		video.EndTimestamp = eventData.EndTimestamp
		video.UUID = eventData.UUID
		video.HoverImgPath = eventData.HoverImgPath
		video.Time = eventData.Time
		video.Hour = eventData.Hour
		jsonb, _ := json.Marshal(video)
		log.Println("json:", string(jsonb))
		err = s.VideoRecordService.Create(video)
		log.Println("EventType4 VideoRecord:", video, err)
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
	openID := c.GetHeader("X-WX-OPENID")
	uuid := c.Param("uuid")
	data, err := s.EventService.GetHighlightsVideos(uuid)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	rs := s.FetchDownloadStatusClips(data, openID)
	c.JSON(200, resp.ToStruct(rs, err))
}

func (s *Service) FetchDownloadStatusClips(datas []model.VideoClips, openId string) []*VideoClipsRes {
	results := make([]*VideoClipsRes, 0)
	for _, v := range datas {
		rr := &VideoClipsRes{}
		rr.ID = v.ID
		rr.FilePath = v.FilePath
		rr.Team = v.Team
		rr.CourtUUID = v.CourtUUID
		rr.HoverImgPath = v.HoverImgPath
		rr.CreateTime = v.CreateTime
		rr.UpdateTime = v.UpdateTime
		rr.VideoType = v.VideoType
		rr.Time = v.Time
		rr.TimeRange = v.TimeRange
		d, err := s.DownloadRecordService.GetByOpenIdResourceIdAndresourceType(openId, v.ID, 30)
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

// 获取视频图册
func (s *Service) GetVideoImg(c *gin.Context) {
	uuid := c.Param("uuid")
	imgTypeStr := c.Param("type")
	imgType, _ := strconv.Atoi(imgTypeStr)
	log.Println("GetVideoImg uuid  imgTypeStr imgType", uuid, imgTypeStr, imgType)
	data, err := s.EventService.GetVideoImg(uuid, int32(imgType))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// 查询时间段
func (s *Service) TimeRange(c *gin.Context) {
	dateStr := c.Query("date")
	date, err := strconv.Atoi(dateStr)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	data, err := s.EventService.GetTimeRange(int32(date))
	if nil == err {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	}
	c.JSON(200, resp.ToStruct(data, err))
}

func (s *Service) TimeRangeV1(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	dateStr := c.Query("date")
	venueIdStr := c.Query("venueId")
	courtIdStr := c.Query("courtId")
	date, err := strconv.Atoi(dateStr)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	venueId, err := strconv.Atoi(venueIdStr)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	courtId, err := strconv.Atoi(courtIdStr)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	court, err := s.CourtService.GetCourtByID(int32(courtId))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	s.VipService.UpdateLastVidCid(openID, court.VenueId, court.ID)
	// data, err := s.EventService.GetTimeRangeV1(int32(date), int32(venueId), court.CourtCode)
	data, err := s.VideoRecordService.GetTimeRangeV1(int32(date), int32(venueId), court.CourtCode)
	if nil == err {
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
	}
	log.Println("data", data)
	res := make([]TimeRangeRes, 0)
	for _, v := range data {
		tr := TimeRangeRes{}
		tr.CourtId = int32(courtId)
		tr.VenueId = int32(venueId)
		tr.Date = int32(date)
		tr.CourtCode = court.CourtCode
		tr.Hour = v
		vs, err := s.EventService.GetVideoList(int32(date), court.CourtCode, v, int32(venueId))
		if err != nil {
			tr.VideoCnt = 0
			continue
		}
		tr.VideoCnt = int32(len(vs))
		cliCnt := 0
		for _, vv := range vs {
			c, err := s.EventService.VideoClipsDao.GetByCourtUuid(vv.UUID)
			if err != nil {
				tr.VideoClipsCnt = 0
				continue
			}
			cliCnt = cliCnt + len(c)
		}
		tr.VideoClipsCnt = int32(cliCnt)
		records, err := s.VideoRecordService.GetVideoRecords(int32(venueId), int32(courtId), int32(date), v)
		if nil != err {
			tr.VideoRecordCnt = 0
		} else {
			tr.VideoRecordCnt = int32(len(records))
		}
		res = append(res, tr)
	}
	c.JSON(200, resp.ToStruct(res, err))
}

func (s *Service) GetVideoList(c *gin.Context) {
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
	court, err := s.CourtService.GetCourtByID(int32(courtId))
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	data, err := s.EventService.GetVideoList(int32(date), court.CourtCode, int32(hour), int32(venueId))
	rs := s.FetchDownloadStatusVideo(data, openID)
	c.JSON(200, resp.ToStruct(rs, err))
}
func (s *Service) FetchDownloadStatusVideo(datas []*model.Video, openId string) []*VideoRes {
	results := make([]*VideoRes, 0)
	for _, v := range datas {
		rr := &VideoRes{}
		rr.UUID = v.UUID
		rr.ID = v.ID
		rr.FilePath = v.FilePath
		rr.Date = v.Date
		rr.Time = v.Time
		rr.Type = v.Type
		rr.VenueId = v.VenueId
		rr.Court = v.Court
		rr.Hour = v.Hour
		rr.FileName = v.FileName
		rr.VideoName = v.VideoName
		rr.FileType = v.FileType
		rr.CreatedTime = v.CreatedTime
		rr.UpdatedTime = v.UpdatedTime
		rr.StartTime = v.StartTime
		rr.EndTime = v.EndTime
		rr.TeamAImgPath = v.TeamAImgPath
		rr.TeamBImgPath = v.TeamBImgPath
		rr.HoverImgPath = v.HoverImgPath
		d, err := s.DownloadRecordService.GetByOpenIdResourceIdAndresourceType(openId, v.ID, 10)
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

func (s *Service) GetVideoDetails(c *gin.Context) {
	uuid := c.Param("uuid")
	data, err := s.EventService.VideoDetail(uuid)
	c.JSON(200, resp.ToStruct(data, err))
}

func (s *Service) GetClipsVideoDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	data, err := s.EventService.GetAiVideoDetail(int32(id))
	c.JSON(200, resp.ToStruct(data, err))
}
