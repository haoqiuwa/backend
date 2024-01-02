package service

import (
	"encoding/json"
	"io"
	"net/http"
	"wxcloudrun-golang/internal/pkg/resp"
	"wxcloudrun-golang/internal/pkg/util"

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
		c.JSON(http.StatusOK, resp.ToStruct(true, nil))
	}
	c.JSON(http.StatusOK, resp.ToStruct(v, nil))
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
	c.JSON(http.StatusOK, resp.ToStruct(nil, err))
}

func (s *Service) UserOpenId(c *gin.Context) {
	openID := c.GetHeader("X-WX-OPENID")
	c.JSON(http.StatusOK, resp.ToStruct(openID, nil))
}

func (s *Service) AccessToken(c *gin.Context) {
	token, err := util.GetAccessToken()
	if nil != err {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, string(token))
}

func (s *Service) QRCode(c *gin.Context) {
	idStr := c.Param("id")
	typeStr := c.Param("type") //正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版。
	req := util.QRCodeReq{
		Scene:      "venueId=" + idStr,
		Page:       "pages/introduce/index",
		CheckPath:  false,
		EnvVersion: typeStr,
	}
	qrCode, err := util.GetUnlimitedQRCode(req)
	if nil != err {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "image/jpeg", qrCode)
}
