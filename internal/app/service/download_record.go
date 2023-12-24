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

type PriceConfig struct {
	Version          string `json:"version"`
	CourtVideoPrice  int32  `json:"court_video_price"`
	VideoRecordPrice int32  `json:"video_record_price"`
	AiClipsPrice     int32  `json:"ai_clips_price"`
}

type Config struct {
	PriceConfig PriceConfig `json:"price_config"`
}

func (s *Service) GetUserDownloadList(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	// rt := c.Query("resource_type")
	// resourceType, err := strconv.Atoi(rt)
	// if nil != err {
	// 	c.JSON(400, "参数错误")
	// 	return
	// }
	page, err := strconv.Atoi(pageStr)
	if nil != err {
		c.JSON(400, "参数错误")
		return
	}
	offset := (int32(page) - 1) * pageSize
	pageInfo := resp.PageInfo{}
	pageInfo.Page = int32(page)
	r, err := s.DownloadRecordService.GetByOpenIdPage(openID, offset, pageSize)
	if nil != err {
		c.JSON(200, "暂无数据")
		return
	}
	if len(r) == int(pageSize) {
		pageInfo.HasMore = true
	}
	pageInfo.PageData = r
	c.JSON(200, resp.ToStruct(pageInfo, err))
}

func (s *Service) GetDownloadRecordById(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if nil != err {
		c.JSON(400, "参数错误")
		return
	}
	r, err := s.DownloadRecordService.GetById(int32(id))
	c.JSON(200, resp.ToStruct(r, err))
}

// 下载
func (s *Service) UserDownload(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	log.Println("UserDownload openId:", openID)
	body, _ := io.ReadAll(c.Request.Body)
	userDownload := &request.UserDownload{}
	err := json.Unmarshal(body, userDownload)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	log.Println("UserDownload req:", userDownload)
	v, err := s.VipService.GetByOpenID(openID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	log.Println("UserDownload vip:", v)
	alreadyDr, err := s.DownloadRecordService.GetByOpenIdResourceIdAndresourceType(openID, userDownload.ResourceId, userDownload.ResourceType)
	//已经下载过了直接返回记录
	if err != nil {
		log.Println("UserDownload GetByOpenIdResourceIdAndresourceType err", err)
	}
	log.Println("UserDownload alreadyDr:", alreadyDr)
	if nil != alreadyDr && alreadyDr.ID > 0 {
		c.JSON(200, resp.ToStruct(alreadyDr, err))
		return
	}
	dr := model.DownloadRecord{}
	config := Config{}
	log.Println("UserDownload type: ", userDownload.ResourceType)
	switch userDownload.ResourceType {
	case 10: //场次回放
		video, err := s.EventService.VideoDao.GetVideoById(userDownload.ResourceId)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		log.Println("UserDownload video: ", video)
		venue, err := s.VenueService.GetVenueById(video.VenueId)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		log.Println("UserDownload venue: ", venue)
		conf := venue.VenueConf
		config := Config{}
		err = json.Unmarshal([]byte(conf), &config)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		log.Println("UserDownload venue conf: ", conf)
		court, err := s.CourtService.GetCourtByID(video.Court)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		log.Println("UserDownload venue court: ", court)
		dr.ResourceType = userDownload.ResourceType
		dr.ResourceUUID = video.UUID
		dr.VenueId = video.VenueId
		dr.CourtId = video.Court
		dr.FilePath = video.FilePath
		dr.HoverImgPath = video.HoverImgPath
		dr.ResourceId = video.ID
		dr.OpenID = openID
		dr.CastDiamond = config.PriceConfig.CourtVideoPrice
		dr.CurrentDiamond = v.Count - dr.CastDiamond
		dr.VenueName = venue.VenueName
		dr.CourtName = court.CourtName
		dr.CreateTime = time.Now()
		dr.UpdateTime = time.Now()
	case 20: //录像
		vr, err := s.VideoRecordService.GetById(userDownload.ResourceId)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		venue, err := s.VenueService.GetVenueById(vr.VenueId)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		conf := venue.VenueConf
		config := Config{}
		err = json.Unmarshal([]byte(conf), &config)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		court, err := s.CourtService.GetCourtByID(vr.CourtId)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		dr.ResourceType = userDownload.ResourceType
		dr.ResourceUUID = vr.UUID
		dr.VenueId = vr.VenueId
		dr.CourtId = vr.CourtId
		dr.FilePath = vr.FilePath
		dr.HoverImgPath = vr.HoverImgPath
		dr.ResourceId = vr.ID
		dr.OpenID = openID
		dr.CastDiamond = config.PriceConfig.VideoRecordPrice
		dr.CurrentDiamond = v.Count - dr.CastDiamond
		dr.VenueName = venue.VenueName
		dr.CourtName = court.CourtName
		dr.CreateTime = time.Now()
		dr.UpdateTime = time.Now()
	case 30: //ai集锦
		clips, err := s.EventService.VideoClipsDao.GetById(userDownload.ResourceId)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		video, err := s.EventService.VideoDao.GetVideoByUUID(clips.CourtUUID)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		venue, err := s.VenueService.GetVenueById(video.VenueId)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		conf := venue.VenueConf
		err = json.Unmarshal([]byte(conf), &config)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		court, err := s.CourtService.GetCourtByID(video.Court)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		dr.ResourceType = userDownload.ResourceType
		dr.ResourceUUID = ""
		dr.VenueId = video.VenueId
		dr.CourtId = video.Court
		dr.FilePath = clips.FilePath
		dr.HoverImgPath = clips.HoverImgPath
		dr.ResourceId = clips.ID
		dr.OpenID = openID
		dr.CastDiamond = config.PriceConfig.AiClipsPrice
		dr.CurrentDiamond = v.Count - dr.CastDiamond
		dr.VenueName = venue.VenueName
		dr.CourtName = court.CourtName
		dr.CreateTime = time.Now()
		dr.UpdateTime = time.Now()
	default:
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
	}
	log.Println("UserDownload  dr: ", dr)
	vp, err := s.VipService.GetByOpenID(openID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if dr.CastDiamond > vp.Count {
		c.JSON(200, resp.Fail(5000, "钻石不足"))
		return
	}
	r, err := s.DownloadRecordService.Create(&dr)
	if err != nil {
		log.Println("DownloadRecordService.Create err", err)
		c.JSON(500, err.Error())
		return
	}
	_, err = s.VipService.UpdateRemainingCount(openID, -r.CastDiamond)
	if err != nil {
		log.Println("VipService.UpdateRemainingCount err", err)
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(dr, err))
}
