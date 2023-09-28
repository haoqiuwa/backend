package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"wxcloudrun-golang/internal/pkg/model"
)

type Service struct {
	UserDao *model.User
}

func NewService() *Service {
	return &Service{}
}

type WXLoginResp struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	DataList []struct {
		CloudID string `json:"cloud_id"`
		JSON    string `json:"json"`
	} `json:"data_list"`
}

type PhoneInfo struct {
	Data struct {
		PhoneNumber string `json:"phoneNumber"`
	}
}

func (s *Service) WXLogin(openid string, cloudID string) (bool, error) {
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url := fmt.Sprintf("http://api.weixin.qq.com/wxa/getopendata?openid=%s", openid)
	// set body
	body, err := json.Marshal(map[string]interface{}{
		"cloudid_list": []string{cloudID},
	})
	// 创建http post请求
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}
	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	if err := json.Unmarshal(respBody, &wxResp); err != nil {
		return false, err
	}
	if wxResp.Errcode != 0 {
		return false, fmt.Errorf("微信登录失败: %s", wxResp.Errmsg)
	}
	// 解析手机号
	phoneInfo := PhoneInfo{}
	if err := json.Unmarshal([]byte(wxResp.DataList[0].JSON), &phoneInfo); err != nil {
		return false, err
	}
	// print data
	if wxResp.DataList != nil && len(wxResp.DataList) > 0 && phoneInfo.Data.PhoneNumber != "" {
		_, err = s.UserDao.Create(&model.User{
			OpenID:      openid,
			Phone:       phoneInfo.Data.PhoneNumber,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (s *Service) StoreCourt(openid string, court int32) error {
	user, err := s.UserDao.Gets(&model.User{OpenID: openid})
	if err != nil {
		return err
	}
	if len(user) == 0 {
		return fmt.Errorf("用户不存在")
	}
	if user[0].Court != 0 {
		return nil
	}
	// update this user
	user[0].Court = court
	user[0].UpdatedTime = time.Now()
	_, err = s.UserDao.Update(&user[0])
	return err
}
