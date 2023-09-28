package tcos

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type TmpAuth struct {
	TmpSecretId  string        `json:"TmpSecretId"`
	TmpSecretKey string        `json:"TmpSecretKey"`
	Token        string        `json:"Token"`
	ExpiredTime  time.Duration `json:"ExpiredTime"`
}

// 拉取cos文件夹下的文件列表
// https://cloud.tencent.com/document/product/436/7743
func GetCosFileList(prefix string) ([]string, error) {
	resp, err := http.Get("http://api.weixin.qq.com/_/cos/getauth")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	// unmarchal resp.Body to TmpAuth
	temAuth := &TmpAuth{}
	err = json.NewDecoder(resp.Body).Decode(temAuth)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	u, _ := url.Parse("https://7072-prod-2gicsblt193f5dc8-1318337180.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: temAuth.TmpSecretId,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey:    temAuth.TmpSecretKey,
			SessionToken: temAuth.Token,
			Expire:       temAuth.ExpiredTime,
		},
	})
	opt := &cos.BucketGetOptions{
		Prefix:  prefix,
		MaxKeys: 1000,
	}
	cos, _, err := client.Bucket.Get(context.Background(), opt)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	var videoIDs []string
	for _, v := range cos.Contents {
		videoIDs = append(videoIDs, fmt.Sprintf("cloud://prod-2gicsblt193f5dc8."+
			"7072-prod-2gicsblt193f5dc8-1318337180/%s", v.Key))
	}
	return videoIDs, nil
}
