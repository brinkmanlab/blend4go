package histories_test

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/histories"
	"testing"
)

func TestHistory_Delete(t *testing.T) {
	type args struct {
		ctx   context.Context
		purge bool
	}
	tests := []struct {
		name    string
		args    args
		success func(*histories.History) bool
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				ctx:   context.Background(),
				purge: false,
			},
			success: func(history *histories.History) bool {
				return history.Deleted && !history.Purged
			},
			wantErr: false,
		}, {
			name: "Purged",
			args: args{
				ctx:   context.Background(),
				purge: true,
			},
			success: func(history *histories.History) bool {
				return history.Deleted && history.Purged
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := histories.NewHistory(tt.args.ctx, galaxyInstance, "test")
			if err != nil {
				t.Fatalf("Failed to create history %v", err)
			}
			if err := h.Delete(tt.args.ctx, tt.args.purge); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.success(h) {
				t.Errorf("Delete() failed to delete history")
			}
		})
	}
}

func TestHistory_GetBasePath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Basic",
			want: histories.BasePath,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &histories.History{}
			if got := h.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistory_GetID(t *testing.T) {
	type fields struct {
		Id blend4go.GalaxyID
	}
	tests := []struct {
		name   string
		fields fields
		want   blend4go.GalaxyID
	}{
		{
			name:   "Basic",
			fields: fields{Id: "test"},
			want:   "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &histories.History{
				Id: tt.fields.Id,
			}
			if got := h.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistory_NewHistory(t *testing.T) {
	type args struct {
		ctx  context.Context
		g    *blend4go.GalaxyInstance
		name string
	}
	tests := []struct {
		name    string
		args    args
		success func(*histories.History) bool
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				ctx:  context.Background(),
				g:    galaxyInstance,
				name: "test",
			},
			success: func(history *histories.History) bool {
				return history.Id != "" &&
					history.Name == "test"
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := histories.NewHistory(tt.args.ctx, tt.args.g, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Cleanup(func() {
				if err := got.Delete(context.Background(), true); err != nil {
					t.Fatalf("Failed to delete history: %v", err)
					return
				}
			})
			if !tt.success(got) {
				t.Errorf("NewHistory() got = %v", got)
			}
		})
	}
}

func TestHistory_SetGalaxyInstance(t *testing.T) {
	t.SkipNow()
}

func TestHistory_SetID(t *testing.T) {
	type fields struct {
		Id blend4go.GalaxyID
	}
	type args struct {
		id blend4go.GalaxyID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		success func(*histories.History) bool
	}{
		{
			name:   "Basic",
			fields: fields{Id: "test"},
			args:   args{id: "newid"},
			success: func(history *histories.History) bool {
				return history.Id == "newid"
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &histories.History{
				Id: tt.fields.Id,
			}
			h.SetID(tt.args.id)
			if !tt.success(h) {
				t.Errorf("SetID failed to set id: %v", h)
			}
		})
	}
}

func TestHistory_Undelete(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		success func(*histories.History) bool
	}{
		{
			name:    "Basic",
			args:    args{ctx: context.Background()},
			wantErr: false,
			success: func(history *histories.History) bool {
				return !history.Deleted
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHistory(t, "test")
			if err := h.Delete(tt.args.ctx, false); err != nil {
				t.Fatalf("Failed to delete history: %v", err)
			}
			if !h.Deleted {
				t.Fatalf("Failed to delete history %v", h)
			}
			if err := h.Undelete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Undelete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.success(h) {
				t.Errorf("Undelete failed to undelete history: %v", h)
			}
		})
	}
}

func TestHistory_Update(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		modify  func(*histories.History)
		success func(*histories.History) bool
		args    args
		wantErr bool
	}{
		{
			name: "Name",
			modify: func(history *histories.History) {
				history.Name = "Name"
			},
			success: func(history *histories.History) bool {
				return history.Name == "Name"
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		}, {
			name: "Tags",
			modify: func(history *histories.History) {
				history.Tags = []string{"test"}
			},
			success: func(history *histories.History) bool {
				return history.Tags[0] == "test"
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		}, {
			name: "Annotation",
			modify: func(history *histories.History) {
				history.Annotation = "Annotation"
			},
			success: func(history *histories.History) bool {
				return history.Annotation == "Annotation"
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		}, {
			name: "Slug",
			modify: func(history *histories.History) {
				history.Slug = "Slug"
			},
			success: func(history *histories.History) bool {
				return history.Slug == "Slug"
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		}, {
			name: "Published",
			modify: func(history *histories.History) {
				history.Published = true
			},
			success: func(history *histories.History) bool {
				return history.Published
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHistory(t, "test")
			if tt.success(h) {
				t.Fatalf("Initial value == expected value: %v", h)
			}
			tt.modify(h)
			if err := h.Update(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.success(h) {
				t.Errorf("Update() got = %v", h)
			}
		})
	}
}
