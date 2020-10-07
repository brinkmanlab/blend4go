package users_test

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/test_util"
	"github.com/brinkmanlab/blend4go/users"
	"testing"
)

var galaxyInstance = test_util.NewTestInstance()

func TestGet(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
		id  blend4go.GalaxyID
	}
	tests := []struct {
		name    string
		args    args
		success func(*users.User) bool
		wantErr bool
	}{
		{
			name: "current",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
				id:  "",
			},
			success: func(user *users.User) bool {
				return user != nil &&
					user.Id != "" &&
					user.Username == test_util.User &&
					user.Email != ""
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := users.Get(tt.args.ctx, tt.args.g, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.success(got) {
				t.Errorf("Get() got = %v", got)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		ctx         context.Context
		g           *blend4go.GalaxyInstance
		deleted     bool
		filterEmail string
		filterName  string
		filterAny   string
	}
	tests := []struct {
		name    string
		args    args
		success func([]*users.User) bool
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				ctx:         context.Background(),
				g:           galaxyInstance,
				deleted:     false,
				filterEmail: "",
				filterName:  "",
				filterAny:   "",
			},
			success: func(u []*users.User) bool {
				return len(u) > 0 && !u[0].Deleted
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := users.List(tt.args.ctx, tt.args.g, tt.args.deleted, tt.args.filterEmail, tt.args.filterName, tt.args.filterAny)
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
