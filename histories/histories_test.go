package histories

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
		id  blend4go.GalaxyID
	}
	tests := []struct {
		name    string
		args    args
		want    *History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.ctx, tt.args.g, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMostRecent(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		want    *History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMostRecent(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMostRecent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMostRecent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPublished(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		want    []History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPublished(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPublished() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPublished() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSharedWithMe(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		want    []History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSharedWithMe(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSharedWithMe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSharedWithMe() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		want    []History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListDeleted(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		want    []History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListDeleted(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDeleted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListDeleted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list(t *testing.T) {
	type args struct {
		ctx      context.Context
		g        *blend4go.GalaxyInstance
		category string
	}
	tests := []struct {
		name    string
		args    args
		want    []History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := list(tt.args.ctx, tt.args.g, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("list() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("list() got = %v, want %v", got, tt.want)
			}
		})
	}
}
