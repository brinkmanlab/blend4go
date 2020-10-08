package histories_test

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/histories"
	"github.com/brinkmanlab/blend4go/test_util"
	"reflect"
	"testing"
)

var galaxyInstance = test_util.NewTestInstance()

func newTestHistory(t *testing.T, name string) *histories.History {
	h, err := histories.NewHistory(context.Background(), galaxyInstance, name)
	if err != nil {
		t.Fatalf("Failed to create history %v", err)
	}
	t.Cleanup(func() {
		if err := h.Delete(context.Background(), true); err != nil {
			t.Fatalf("Failed to delete history: %v", err)
		}
	})
	return h
}

func TestGet(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
		id  blend4go.GalaxyID
	}
	tests := []struct {
		name    string
		args    args
		success func(got *histories.History, want *histories.History) bool
		want    *histories.History
		wantErr bool
	}{
		{
			name: "None",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
				id:  "test",
			},
			success: func(got *histories.History, want *histories.History) bool {
				return got == nil
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Basic",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
				id:  "",
			},
			success: func(got *histories.History, want *histories.History) bool {
				return got != nil && want != nil && got.Id == want.Id
			},
			want:    newTestHistory(t, "test"),
			wantErr: false,
		}, {
			name: "Missing id",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
				id:  "",
			},
			success: func(got *histories.History, want *histories.History) bool {
				return true
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.args.id
			if id == "" && tt.want != nil {
				id = tt.want.Id
			}
			got, err := histories.Get(tt.args.ctx, tt.args.g, id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.success(got, tt.want) {
				t.Errorf("Get() got = %v, want = %v", got, tt.want)
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
		want    *histories.History
		wantErr bool
	}{
		{
			name: "None",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Basic",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
			},
			want:    newTestHistory(t, "test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := histories.GetMostRecent(tt.args.ctx, tt.args.g)
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
		prepare func(*testing.T)
		success func([]*histories.History) bool
		wantErr bool
	}{
		{
			name: "None",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
			},
			prepare: func(t *testing.T) {},
			success: func(i []*histories.History) bool {
				return len(i) == 0
			},
			wantErr: false,
		}, {
			name: "Basic",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
			},
			prepare: func(t *testing.T) {
				h := newTestHistory(t, "test1")
				h.Published = true
				if err := h.Update(context.Background()); err != nil {
					t.Fatalf("Failed to prepare published history: %v", err)
				}
			},
			success: func(i []*histories.History) bool {
				return len(i) > 0
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(t)
			got, err := histories.GetPublished(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPublished() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.success(got) {
				t.Errorf("GetPublished() got = %v", got)
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
		prepare func(*testing.T)
		success func([]*histories.History) bool
		wantErr bool
	}{
		{
			name: "None",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
			},
			prepare: func(t *testing.T) {},
			success: func(i []*histories.History) bool {
				return len(i) == 0
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(t)
			got, err := histories.GetSharedWithMe(tt.args.ctx, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSharedWithMe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.success(got) {
				t.Errorf("GetSharedWithMe() got = %v", got)
			}
		})
	}
}

func testList(t *testing.T, f func(ctx context.Context, g *blend4go.GalaxyInstance) ([]*histories.History, error)) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		prepare func(*testing.T)
		success func([]*histories.History) bool
		wantErr bool
	}{
		{
			name: "None",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
			},
			prepare: func(t *testing.T) {},
			success: func(i []*histories.History) bool {
				return len(i) == 0
			},
			wantErr: false,
		}, {
			name: "Basic",
			args: args{
				ctx: context.Background(),
				g:   galaxyInstance,
			},
			prepare: func(t *testing.T) {
				newTestHistory(t, "test")
			},
			success: func(i []*histories.History) bool {
				return len(i) == 1
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(t)
			got, err := f(tt.args.ctx, tt.args.g)
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

func TestList(t *testing.T) {
	testList(t, histories.List)
}

func TestListDeleted(t *testing.T) {
	testList(t, histories.ListDeleted)
}
