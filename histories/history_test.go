package histories

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"reflect"
	"testing"
)

func TestHistory_Delete(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
			if err := h.Delete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHistory_GetBasePath(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
			if got := h.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistory_GetID(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	tests := []struct {
		name   string
		fields fields
		want   blend4go.GalaxyID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
			if got := h.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistory_NewHistory(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	type args struct {
		ctx  context.Context
		g    *blend4go.GalaxyInstance
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *History
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
			got, err := h.NewHistory(tt.args.ctx, tt.args.g, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHistory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistory_SetGalaxyInstance(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	type args struct {
		g *blend4go.GalaxyInstance
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
		})
	}
}

func TestHistory_SetID(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	type args struct {
		id blend4go.GalaxyID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
		})
	}
}

func TestHistory_Undelete(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
			if err := h.Undelete(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Undelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHistory_Update(t *testing.T) {
	type fields struct {
		galaxyInstance  *blend4go.GalaxyInstance
		Id              blend4go.GalaxyID
		Importable      bool
		CreateTime      string
		ContentsUrl     string
		Size            uint
		UserId          blend4go.GalaxyID
		UsernameAndSlug string
		Annotation      string
		StateDetails    map[string]uint
		State           string
		Empty           bool
		UpdateTime      string
		Tags            []string
		Deleted         bool
		GenomeBuild     string
		Slug            string
		Name            string
		Url             string
		StateIds        map[string][]blend4go.GalaxyID
		Published       bool
		ModelClass      string
		Purged          bool
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &History{
				galaxyInstance:  tt.fields.galaxyInstance,
				Id:              tt.fields.Id,
				Importable:      tt.fields.Importable,
				CreateTime:      tt.fields.CreateTime,
				ContentsUrl:     tt.fields.ContentsUrl,
				Size:            tt.fields.Size,
				UserId:          tt.fields.UserId,
				UsernameAndSlug: tt.fields.UsernameAndSlug,
				Annotation:      tt.fields.Annotation,
				StateDetails:    tt.fields.StateDetails,
				State:           tt.fields.State,
				Empty:           tt.fields.Empty,
				UpdateTime:      tt.fields.UpdateTime,
				Tags:            tt.fields.Tags,
				Deleted:         tt.fields.Deleted,
				GenomeBuild:     tt.fields.GenomeBuild,
				Slug:            tt.fields.Slug,
				Name:            tt.fields.Name,
				Url:             tt.fields.Url,
				StateIds:        tt.fields.StateIds,
				Published:       tt.fields.Published,
				ModelClass:      tt.fields.ModelClass,
				Purged:          tt.fields.Purged,
			}
			if err := h.Update(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
