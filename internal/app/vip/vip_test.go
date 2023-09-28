package vip

import (
	"os"
	"testing"
	"wxcloudrun-golang/internal/pkg/db"
)

var s = NewService()

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}
func TestService_GetRemainingCount(t *testing.T) {
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
			args:    args{openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetRemainingCount(tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemainingCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_UpdateRemainingCount(t *testing.T) {
	type args struct {
		openID     string
		countToAdd int32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc", countToAdd: -1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateRemainingCount(tt.args.openID, tt.args.countToAdd)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRemainingCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_CreateOrder(t *testing.T) {
	type args struct {
		openID    string
		orderType int32
		cost      float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc", orderType: 1, cost: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.CreateOrder(tt.args.openID, tt.args.orderType, tt.args.cost)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetOrdersByOpenID(t *testing.T) {
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
			args:    args{openID: "oueu25X3eun7K9zJ6UpCUQiEO0yc"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.GetOrdersByOpenID(tt.args.openID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrdersByOpenID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
