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

	router.GET("/events", service.GetEventVideos)
	router.POST("/events", service.StartEvent)
	router.DELETE("/events/:id", service.DeleteEvent)
	router.POST("/collects", service.ToggleCollectVideo)
	router.GET("/user/collects", service.GetCollectVideos)

	router.GET("/user/events/:id", service.GetEventInfo)
	router.GET("/recommend/videos", service.GetRecommendVideos)

	log.Fatal(router.Run())
}
