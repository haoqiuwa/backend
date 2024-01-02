package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

type QRCodeReq struct {
	Scene      string `json:"scene"`
	Page       string `json:"page"`
	CheckPath  bool   `json:"check_path"`
	EnvVersion string `json:"env_version"` //正式版为 "release"，体验版为 "trial"，开发版为 "develop"。默认是正式版。
}

func GetAccessToken() (string, error) {
	appId := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, appSecret)
	log.Println("accessToken url==>>", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("accessToken body==>>", string(body))
	var accessToken AccessToken
	// 解析 JSON
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return "", err
	}
	return accessToken.AccessToken, nil
}

func GetUnlimitedQRCode(req QRCodeReq) ([]byte, error) {
	// 获取 Access Token
	accessToken, err := GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 构建获取小程序码的接口URL
	apiURL := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", accessToken)
	// 将结构体转换为 JSON 字符串
	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	// 发送 HTTP POST 请求获取小程序码
	response, err := http.Post(apiURL, "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	// 读取响应内容
	qrCode, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}
