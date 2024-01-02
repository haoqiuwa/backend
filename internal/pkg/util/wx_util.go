package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int32  `json:"expires_in"`
}

type QRCodeReq struct {
	Scene      string `json:"scene"`
	Page       string `json:"page"`
	CheckPath  bool   `json:"check_path"`
	EnvVersion string `json:"env_version"`
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
