package service

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"wxcloudrun-golang/internal/app/collect"
	"wxcloudrun-golang/internal/app/court"
	"wxcloudrun-golang/internal/app/event"
	"wxcloudrun-golang/internal/app/recommend"
	"wxcloudrun-golang/internal/app/user"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Service struct {
	UserService      *user.Service
	CourtService     *court.Service
	EventService     *event.Service
	CollectService   *collect.Service
	RecommendService *recommend.Service
}

func NewService() *Service {
	return &Service{
		UserService:      user.NewService(),
		CourtService:     court.NewService(),
		EventService:     event.NewService(),
		CollectService:   collect.NewService(),
		RecommendService: recommend.NewService(),
	}
}

// WeChatLogin /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
func (s *Service) WeChatLogin(c *gin.Context) {
	code := c.Query("code") //  获取code
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := s.UserService.WXLogin(code)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	// 保存登录态
	session := sessions.Default(c)
	session.Set("openid", wxLoginResp.OpenId)
	session.Set("sessionKey", wxLoginResp.SessionKey)
	// 这里用openid和sessionkey的串接 进行MD5之后作为该用户的自定义登录态
	mySession := user.GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库中,
	// 但这里要保证mySession 唯一, 以便于用mySession去索引openid 和sessionkey
	c.String(200, mySession)
}

func (s *Service) StartEvent(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	body, _ := ioutil.ReadAll(c.Request.Body)
	newEvent := &model.Event{}
	err := json.Unmarshal(body, newEvent)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	newEvent, err = s.EventService.CreateEvent(openID, newEvent.CourtID, newEvent.Date, newEvent.StartTime,
		newEvent.EndTime)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(newEvent, err))
}

func (s *Service) DeleteEvent(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	eventID := c.Param("id")
	eventIDInt, _ := strconv.Atoi(eventID)
	err := s.EventService.DeleteEvent(openID, int32(eventIDInt))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(nil, err))
}

// 主页面相关

// ToggleCollectVideo 收藏视频
func (s *Service) ToggleCollectVideo(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	newCollect := &model.Collect{}
	err := json.Unmarshal(body, newCollect)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	collectRecord, err := s.CollectService.ToggleCollectVideo(openID, newCollect.FileID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(collectRecord, err))
}

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

// GetEvent 获取用户所属事件的视频
func (s *Service) GetEvent(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	results, err := s.EventService.GetEvents(openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(results, err))
}

// GetEventInfo 获取事件
func (s *Service) GetEventInfo(c *gin.Context) {
	eventID := c.Param("id")
	eventIDInt, _ := strconv.Atoi(eventID)
	event, err := s.EventService.GetEventInfo(int32(eventIDInt))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(event, err))
}

// GetCollectVideos 获取用户收藏的视频
func (s *Service) GetCollectVideos(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	collects, err := s.CollectService.GetCollectByUser(openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(collects, err))
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
