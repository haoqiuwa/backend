package service

import (
	"encoding/json"
	"io"
	"log"
	"strconv"
	"wxcloudrun-golang/internal/pkg/model"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

const pageSize int32 = 200

// ToggleCollectVideo 收藏视频
func (s *Service) ToggleCollectVideo(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	newCollect := &model.Collect{}
	err := json.Unmarshal(body, newCollect)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	collectRecord, err := s.CollectService.ToggleCollectVideo(openID, newCollect.FileID, newCollect.PicURL, newCollect.VideoType)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(collectRecord, err))
}

// GetUserDownload 获取用户下载次数
func (s *Service) GetUserDownload(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	data, err := s.CollectService.GetUserDownload(openID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// GetUserDownloadStatus 获取用户下载状态
func (s *Service) GetUserDownloadStatus(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	videoID := c.Query("file_id")
	data, err := s.CollectService.GetUserDownloadStatus(openID, videoID)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// GetUserDownloads 获取用户下载记录
func (s *Service) GetUserDownloads(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	queryType := c.Query("query_type") // 前端传递
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	if queryType == "" {
		queryType = "video"
	}
	pageInt, _ := strconv.Atoi(page)
	offset := (int32(pageInt) - 1) * pageSize
	data, err := s.CollectService.GetUserDownloads(openID, queryType, offset, pageSize)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// CollectSurvey 下载问卷记录
func (s *Service) CollectSurvey(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	data, err := s.CollectService.CreateSurvey(openID, string(body))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// CollectUserEvent 下载视频记录
func (s *Service) CollectUserEvent(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := io.ReadAll(c.Request.Body)
	userEvent := &model.UserEvent{}
	err := json.Unmarshal(body, userEvent)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	log.Println("CollectUserEvent data:", string(body))
	data, err := s.CollectService.CollectUserEvent(openID, userEvent.FileID, userEvent.EventType, userEvent.FromPage,
		userEvent.VideoType)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(data, err))
}

// GetCollectVideos 获取用户收藏的视频
func (s *Service) GetCollectVideos(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	videoType := c.Query("video_type")
	videoTypeInt, _ := strconv.Atoi(videoType)
	collects, err := s.CollectService.GetCollectByUser(openID, int32(videoTypeInt))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(collects, err))
}
