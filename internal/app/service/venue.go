package service

import (
	"log"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetVenues(c *gin.Context) {
	r, err := s.VenueService.GetVenues()
	log.Println("GetVenues r", r)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(r, err))
}
