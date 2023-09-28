package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"wxcloudrun-golang/internal/pkg/resp"
)

type PhoneReq struct {
	CloudID string `json:"cloud_id"`
}

// WeChatLogin /wechat/applet_login?code=xxx [get]  路由
// 微信小程序登录
func (s *Service) WeChatLogin(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	var phoneReq PhoneReq
	body, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &phoneReq)
	// 根据code获取 openID 和 session_key
	wxLoginResp, err := s.UserService.WXLogin(openID, phoneReq.CloudID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, resp.ToStruct(wxLoginResp, err))
}

type courtReq struct {
	Court int32 `json:"court"`
}

// StoreCourt
func (s *Service) StoreCourt(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	if openID == "" {
		c.JSON(400, "请先登录")
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	var courtReq courtReq
	_ = json.Unmarshal(body, &courtReq)
	err := s.UserService.StoreCourt(openID, courtReq.Court)
	c.JSON(200, resp.ToStruct(nil, err))
}
