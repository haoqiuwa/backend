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
	router.GET("/auth/login", service.WeChatLogin)
	router.GET("/courts", service.GetCounts)
	router.GET("/courts/:id", service.GetCountInfo)
	router.GET("/courts/:id/judge", service.JudgeLocation)

	router.GET("/events", service.GetEvents)
	router.GET("/videos", service.GetEventInfo)
	router.POST("/collects", service.ToggleCollectVideo)
	router.GET("/user/collects", service.GetCollectVideos)

	router.GET("/recommend/videos", service.GetRecommendVideos)

	// 8080 port
	log.Fatal(router.Run())
}
