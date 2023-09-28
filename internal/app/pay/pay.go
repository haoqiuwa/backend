package pay

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var mchID = "1645284099"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

type RequestBody struct {
	Text string `json:"text"`
	Noid string `json:"noid"`
	Fee  int    `json:"fee"`
}

type PayReq struct {
	Body           string `json:"body"`
	OutTradeNo     string `json:"out_trade_no"`
	SubMchID       string `json:"sub_mch_id"`
	TotalFee       int    `json:"total_fee"`
	OpenID         string `json:"openid"`
	SpbillCreateIP string `json:"spbill_create_ip"`
	EnvID          string `json:"env_id"`
	CallbackType   int    `json:"callback_type"`
	Container      struct {
		Service string `json:"service"`
		Path    string `json:"path"`
	} `json:"container"`
}

type Response struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	RespData struct {
		ReturnCode string `json:"return_code"`
		ReturnMsg  string `json:"return_msg"`
		AppID      string `json:"appid"`
		MchID      string `json:"mch_id"`
		SubAppID   string `json:"sub_appid"`
		SubMchID   string `json:"sub_mch_id"`
		NonceStr   string `json:"nonce_str"`
		Sign       string `json:"sign"`
		ResultCode string `json:"result_code"`
		TradeType  string `json:"trade_type"`
		PrepayID   string `json:"prepay_id"`
		Payment    struct {
			AppID     string `json:"appId"`
			TimeStamp string `json:"timeStamp"`
			NonceStr  string `json:"nonceStr"`
			Package   string `json:"package"`
			SignType  string `json:"signType"`
			PaySign   string `json:"paySign"`
		} `json:"payment"`
	} `json:"respdata"`
}

func (s *Service) UnifiedOrder(openID string, ip string, body string) (*Response,
	error) {
	var req RequestBody
	err := json.Unmarshal([]byte(body), &req)
	if err != nil {
		return nil, err
	}
	payReq := PayReq{
		Body:           req.Text,
		OutTradeNo:     req.Noid,
		SubMchID:       mchID, // replace with your merchant ID
		TotalFee:       req.Fee,
		OpenID:         openID,
		SpbillCreateIP: ip,
		EnvID:          "prod-2gicsblt193f5dc8",
		CallbackType:   2,
		Container: struct {
			Service string `json:"service"`
			Path    string `json:"path"`
		}{
			Service: "golang-m7vn-065",
			Path:    "/",
		},
	}
	resp, err := callPay("unifiedOrder", payReq)
	if err != nil {
		return nil, err
	}
	var info *Response
	if err := json.Unmarshal(resp, &info); err != nil {
		return nil, err
	}
	return info, nil
}

func callPay(action string, payBody interface{}) ([]byte, error) {
	url := fmt.Sprintf("http://api.weixin.qq.com/_/pay/%s", action)
	body, err := json.Marshal(payBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
