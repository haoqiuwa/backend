package tcos

import "testing"

func TestGetCosFileList(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				prefix: "highlight/court1/20230103/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vs, _ := GetCosFileList(tt.args.prefix)
			// if nil != err {
			// 	t.Errorf("GetCosFileList() error = %v", err)
			// }
			for v := range vs {
				t.Log("find video:", v)
			}
		})
	}
}
