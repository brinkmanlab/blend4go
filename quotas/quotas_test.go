package quotas

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/test_util"
	"reflect"
	"testing"
)

var galaxyInstance = test_util.NewTestInstance()

func createTestQuota() (test_quota *Quota, err error) {
	test_quota, err = NewQuota(context.Background(), galaxyInstance, "test", "0", "test", SetTo, nil, nil, NotDefault)
	if err != nil {
		test_quota, err = GetName(context.Background(), galaxyInstance, "test")
	}
	if test_quota != nil && test_quota.Deleted {
		err = test_quota.Undelete(context.Background())
	}
	return
}

func TestGet(t *testing.T) {
	test_quota, err := createTestQuota()
	if test_quota == nil {
		t.Fatalf("test quota could not be set up: %v", err)
	}
	defer test_quota.Delete(context.Background(), true)
	type args struct {
		ctx     context.Context
		g       *blend4go.GalaxyInstance
		id      blend4go.GalaxyID
		deleted bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Quota
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				ctx:     context.Background(),
				g:       galaxyInstance,
				id:      test_quota.Id,
				deleted: false,
			},
			want:    test_quota,
			wantErr: false,
		}, {
			name: "not_deleted",
			args: args{
				ctx:     context.Background(),
				g:       galaxyInstance,
				id:      test_quota.Id,
				deleted: true,
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "not_exist",
			args: args{
				ctx:     context.Background(),
				g:       galaxyInstance,
				id:      "fake",
				deleted: false,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.ctx, tt.args.g, tt.args.id, tt.args.deleted)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get()\ngot = %#v\nwant %#v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	test_quota, err := createTestQuota()
	if test_quota == nil {
		t.Fatalf("test quota could not be set up: %v", err)
	}
	defer test_quota.Delete(context.Background(), true)
	type args struct {
		ctx     context.Context
		g       *blend4go.GalaxyInstance
		deleted bool
	}
	tests := []struct {
		name    string
		args    args
		success func([]*Quota) bool
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				ctx:     context.Background(),
				g:       galaxyInstance,
				deleted: false,
			},
			success: func(quotas []*Quota) bool {
				for _, quota := range quotas {
					if quota.Deleted == true {
						return false
					}
				}
				return true
			},
			wantErr: false,
		}, {
			name: "deleted",
			args: args{
				ctx:     context.Background(),
				g:       galaxyInstance,
				deleted: true,
			},
			success: func(quotas []*Quota) bool {
				for _, quota := range quotas {
					if quota.Deleted != true {
						return false
					}
				}
				return true
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.ctx, tt.args.g, tt.args.deleted)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.success(got) {
				t.Errorf("List() got = %v", got)
			}
		})
	}
}

func TestGetName(t *testing.T) {
	test_quota, err := createTestQuota()
	if test_quota == nil {
		t.Fatalf("test quota could not be set up: %v", err)
	}
	defer test_quota.Delete(context.Background(), true)
	type args struct {
		ctx  context.Context
		g    *blend4go.GalaxyInstance
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *Quota
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				ctx:  context.Background(),
				g:    galaxyInstance,
				name: "test",
			},
			want:    test_quota,
			wantErr: false,
		}, {
			name: "not_exist",
			args: args{
				ctx:  context.Background(),
				g:    galaxyInstance,
				name: "foo_test",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetName(tt.args.ctx, tt.args.g, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetName()\ngot = %#v\nwant %#v", got, tt.want)
			}
		})
	}
}
