package event

import (
	"os"
	"testing"
	"wxcloudrun-golang/internal/pkg/db"
)

var s Service

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}

func TestService_GetEventVideos(t *testing.T) {
	type args struct {
		openID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{openID: "ogDJL5R996RQKkZm0VtFaK2-i3ik"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetEvents(tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetEvents(t *testing.T) {
	type args struct {
		courtID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{courtID: "10"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetEvents(tt.args.courtID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetEventInfo(t *testing.T) {
	type args struct {
		courtID string
		hour    int
		openID  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{courtID: "10", hour: 12, openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetEventInfo(tt.args.courtID, tt.args.hour, tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
