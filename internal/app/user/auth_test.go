package user

import (
	"encoding/json"
	"os"
	"testing"
	"wxcloudrun-golang/internal/pkg/db"
)

var s Service

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}

func TestService_WXLogin(t *testing.T) {
	type args struct {
		openid  string
		cloudID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{openid: "oueu25X3eun7K9zJ6UpCUQiEO0yc-i3ik", cloudID: "69_9GwMsLPtiQO8PS5NBc9OJE3swDOLMVCc_7PNZq3q62jxQF4k3n0vTsJfxi8"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.WXLogin(tt.args.openid, tt.args.cloudID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.WXLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetUserPhone(t *testing.T) {
	str := `
{"errcode":0,"errmsg":"ok","data_list":[{"cloud_id":"69_hLOlCnc0ymeJeZXwDn9yBJG2KLZ2fFYuG4I6ZTE0kXV28JZaEVOM9FZB9uk","json":"{ \"cloudID\":\"69_hLOlCnc0ymeJeZXwDn9yBJG2KLZ2fFYuG4I6ZTE0kXV28JZaEVOM9FZB9uk\", \"data\":{\"phoneNumber\":\"13772065985\",\"purePhoneNumber\":\"13772065985\",\"countryCode\":\"86\",\"watermark\":{\"timestamp\":1686993965,\"appid\":\"wx69e45cc989de661d\"}} }"}]}
`
	var resp WXLoginResp
	err := json.Unmarshal([]byte(str), &resp)
	if err != nil {
		t.Error(err)
	}
	var phone PhoneInfo
	err = json.Unmarshal([]byte(resp.DataList[0].JSON), &phone)
	if err != nil {
		t.Error(err)
	}
	print(phone.Data.PhoneNumber)

}

func TestService_StoreCourt(t *testing.T) {
	type args struct {
		openid string
		court  int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{openid: "oueu25X3eun7K9zJ6UpCUQiEO0yc", court: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.StoreCourt(tt.args.openid, tt.args.court); (err != nil) != tt.wantErr {
				t.Errorf("StoreCourt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
