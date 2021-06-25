package dml

import (
	"owen2020/app/models"
	"reflect"
	"testing"
)

func Test_stateDAO_GetMulti(t *testing.T) {
	type args struct {
		idList []int
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]*models.State
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				[]int{1, 2, 3, 4},
			},
			want:    map[int]*models.State{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stateDAO{}
			got, err := s.GetMulti(tt.args.idList)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMulti() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMulti() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_stateDAO_newEngine(t *testing.T) {
//	type args struct {
//		isMaster bool
//	}
//	tests := []struct {
//		name string
//		args args
//		want *gorm.DB
//	}{
//		{
//			name: "master",
//			args: args{isMaster: true},
//			want: &gorm.DB{},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &stateDAO{}
//			if got := s.newEngine(tt.args.isMaster); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("newEngine() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
