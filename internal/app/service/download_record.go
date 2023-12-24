package service

import (
	"encoding/json"
	"io"
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
	// rt := c.Query("resource_type")
	// resourceType, err := strconv.Atoi(rt)
	// if nil != err {
	// 	c.JSON(400, "参数错误")
	// 	return
	// }
	r, err := s.DownloadRecordService.GetByOpenId(openID)
	if nil != err {
		c.JSON(200, "暂无数据")
		return
	}
	c.JSON(200, resp.ToStruct(r, err))
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

func (s *Service) UserDownload(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	userDownload := &request.UserDownload{}
	err := json.Unmarshal(body, userDownload)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	v, err := s.VipService.GetByOpenID(openID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	dr := model.DownloadRecord{}
	config := Config{}
	switch userDownload.ResourceType {
	case 10: //场次回放
		video, err := s.EventService.VideoDao.GetVideoById(userDownload.ResourceId)
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
		config := Config{}
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
		dr.ResourceUUID = video.UUID
		dr.VenueId = video.VenueId
		dr.CourtId = video.Court
		dr.FilePath = video.FilePath
		dr.HoverImgPath = video.HoverImgPath
		dr.ResourceId = video.ID
		dr.OpenID = openID
		dr.CastDiamond = config.PriceConfig.VideoRecordPrice
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
		dr.CastDiamond = config.PriceConfig.VideoRecordPrice
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
	vp, err := s.VipService.GetByOpenID(openID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	if config.PriceConfig.VideoRecordPrice > vp.Count {
		c.JSON(500, "钻石不足")
		return
	}
	r, err := s.DownloadRecordService.Create(&dr)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	_, err = s.VipService.UpdateRemainingCount(openID, -r.CastDiamond)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(dr, err))
}
