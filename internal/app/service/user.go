package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"wxcloudrun-golang/internal/pkg/resp"

	"github.com/gin-gonic/gin"
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
	// var phoneReq PhoneReq
	// body, _ := io.ReadAll(c.Request.Body)
	// _ = json.Unmarshal(body, &phoneReq)
	// // 根据code获取 openID 和 session_key
	// wxLoginResp, err := s.UserService.WXLogin(openID, phoneReq.CloudID)
	// if err != nil {
	// 	c.JSON(400, err.Error())
	// 	return
	// }
	v, err := s.VipService.GetByOpenID(openID)
	if nil != err {
		c.JSON(200, resp.ToStruct(true, nil))
	}
	c.JSON(200, resp.ToStruct(v, nil))
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
	body, _ := io.ReadAll(c.Request.Body)
	var courtReq courtReq
	_ = json.Unmarshal(body, &courtReq)
	err := s.UserService.StoreCourt(openID, courtReq.Court)
	c.JSON(200, resp.ToStruct(nil, err))
}

func (s *Service) UserOpenId(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	c.JSON(200, resp.ToStruct(openID, nil))
}

func (s *Service) AccessToken(c *gin.Context) {
	appId := c.Query("appId")
	appSecret := c.Query("appSecret")
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	log.Println("accessToken url==>>", url)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, err.Error())
	}
	log.Println("accessToken body==>>", string(body))
	c.JSON(200, body)
}
