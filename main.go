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

	// 用户管理
	router.GET("/user/collects", service.GetCollectVideos)
	router.GET("/user/download", service.GetUserDownload)
	router.GET("/user/downloads", service.GetUserDownloads)
	router.POST("/survey", service.CollectSurvey)
	router.GET("/user/download_status", service.GetUserDownloadStatus)

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

	// 视频处理
	router.POST("/videos", service.StoreVideo)
	router.POST("/collects", service.ToggleCollectVideo)
	router.POST("/user/event", service.CollectUserEvent)

	// 暂未启用 场地管理
	router.GET("/courts", service.GetCounts)
	router.GET("/courts/:id", service.GetCountInfo)
	router.GET("/courts/:id/judge", service.JudgeLocation)
	router.GET("/recommend/videos", service.GetRecommendVideos)
	// 8080 port
	log.Fatal(router.Run())
}
