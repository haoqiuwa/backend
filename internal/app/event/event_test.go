package event

import (
	"os"
	"testing"
	"wxcloudrun-golang/internal/pkg/db"
	"wxcloudrun-golang/internal/pkg/model"
)

var s Service

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}

func TestService_GetEvents(t *testing.T) {
	type args struct {
		courtID string
		date    int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{courtID: "10", date: 20230729},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetEvents(tt.args.courtID, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetVideos(t *testing.T) {
	type args struct {
		courtID int32
		hour    int32
		openID  string
		date    int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{courtID: 10, hour: 16, openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc", date: 20230803},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetVideos(tt.args.date, tt.args.courtID, tt.args.hour, tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetRecords(t *testing.T) {
	type args struct {
		courtID int32
		hour    int32
		openID  string
		date    int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{courtID: 10, hour: 11, openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc", date: 20230802},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetRecord(tt.args.date, tt.args.courtID, tt.args.hour, tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetMatchHighlights(t *testing.T) {
	type args struct {
		courtID int32
		hour    int32
		openID  string
		date    int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{courtID: 1, hour: 16, openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc", date: 20230902},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetMatchHighlights(tt.args.date, tt.args.courtID, tt.args.hour, tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
func TestService_StoreVideo(t *testing.T) {
	type args struct {
		video *model.Video
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{video: &model.Video{
				Date:     20230715,
				Hour:     9,
				FileName: "P9-40.MP4",
				Type:     2,
				Court:    10,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.StoreVideo(tt.args.video)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetAIContent(t *testing.T) {
	type args struct {
		date    int32
		courtID int32
		hour    int32
		openID  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				date:    20230920,
				courtID: 10,
				hour:    21,
				openID:  "oueu25X3eun7K9zJ6UpCUQiEO0yc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetAIContent(tt.args.date, tt.args.courtID, tt.args.hour, tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAIContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
