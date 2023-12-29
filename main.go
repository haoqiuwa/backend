package main

import (
	"fmt"
	"log"
	"wxcloudrun-golang/internal/app/service"
	"wxcloudrun-golang/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	service := service.NewService()
	router := gin.Default()
	// 用户信息
	router.POST("/auth/login", service.WeChatLogin)
	router.POST("/user/court", service.StoreCourt)
	router.POST("/user/openid", service.UserOpenId)

	// 用户管理
	router.GET("/user/collects", service.GetCollectVideos)
	router.GET("/user/download", service.GetUserDownload)   //查看有无下载次数
	router.GET("/user/downloads", service.GetUserDownloads) // 查看下载列表
	router.POST("/survey", service.CollectSurvey)
	router.GET("/user/download_status", service.GetUserDownloadStatus) //查看有无下载记录根据field_id

	// 付费相关
	router.GET("/vip/count", service.GetVipCount)
	router.GET("/vip/orders", service.GetVipOrders)
	router.POST("/vip/orders", service.CreateVipOrder)
	router.POST("/vip/count", service.UpdateVipCount)
	router.POST("/vip/pay", service.UnifiedOrder)

	// 视频页面
	router.GET("/events", service.GetEvents)
	router.GET("/videos", service.GetVideos)
	router.GET("/records", service.GetRecords)
	router.GET("/match/highlights", service.GetMatchHighlights)
	router.GET("/match/records", service.GetMatchRecords)
	router.GET("/aigc/contents", service.GetAIContents)
	//时间段筛选
	router.GET("/time/range", service.TimeRange)
	router.GET("/time/range/v1", service.TimeRangeV1)

	//场次ai视频
	router.GET("/ai/videos/:uuid", service.GetAiVideos)
	//场次集锦视频
	router.GET("/highlights/videos/:uuid", service.GetHighlightsVideos)
	//场次图片
	router.GET("/videos/img/:uuid/:type", service.GetVideoImg)
	//时间段筛选回放视频列表
	router.GET("/video/list", service.GetVideoList)
	router.GET("/video/recods", service.GetVideoRecords)
	//场次回放视频
	router.GET("/video/detail/:uuid", service.GetVideoDetails)
	//ai/集锦视频详情
	router.GET("/clips/video/detail/:id", service.GetClipsVideoDetail)

	// 视频处理
	router.POST("/videos", service.StoreVideo)
	router.POST("/videos/event/v1", service.HandlePushEvent)
	router.POST("/collects", service.ToggleCollectVideo)
	router.POST("/user/event", service.CollectUserEvent)

	//  场地管理
	router.GET("/venues", service.GetVenues)
	router.GET("/venue/detail/:id", service.GetVenueDetail)
	router.GET("/courts", service.GetCounts)
	// 活动
	router.GET("/diamond/activity", service.GetActivity)
	router.POST("/use/diamond/activity", service.UseActivity)
	//下载
	router.GET("/downlaod/record/list", service.GetUserDownloadList)
	router.GET("/downlaod/record", service.GetDownloadRecordById)
	router.POST("/user/downlaod/record", service.UserDownload)
	// 8080 port
	log.Fatal(router.Run())
}
