package repositories

import (
	"context"
	"github.com/brinkmanlab/blend4go"
	"reflect"
	"testing"
)

func TestCheckForUpdates(t *testing.T) {
	t.SkipNow()
	return
	type args struct {
		ctx    context.Context
		g      *blend4go.GalaxyInstance
		repoID blend4go.GalaxyID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckForUpdates(tt.args.ctx, tt.args.g, tt.args.repoID); (err != nil) != tt.wantErr {
				t.Errorf("CheckForUpdates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
		want    *Repository
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

func TestInstall(t *testing.T) {
	type args struct {
		ctx                           context.Context
		g                             *blend4go.GalaxyInstance
		toolShedUrl                   string
		name                          string
		owner                         string
		changesetRevision             string
		installToolDependencies       bool
		installRepositoryDependencies bool
		installResolverDependencies   bool
		toolPanelSectionId            blend4go.GalaxyID
		newToolPanelSectionLabel      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Install(tt.args.ctx, tt.args.g, tt.args.toolShedUrl, tt.args.name, tt.args.owner, tt.args.changesetRevision, tt.args.installToolDependencies, tt.args.installRepositoryDependencies, tt.args.installResolverDependencies, tt.args.toolPanelSectionId, tt.args.newToolPanelSectionLabel); (err != nil) != tt.wantErr {
				t.Errorf("Install() error = %v, wantErr %v", err, tt.wantErr)
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
		want    []*Repository
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

func TestResetMetadataAll(t *testing.T) {
	type args struct {
		ctx context.Context
		g   *blend4go.GalaxyInstance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ResetMetadataAll(tt.args.ctx, tt.args.g); (err != nil) != tt.wantErr {
				t.Errorf("ResetMetadataAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUninstall(t *testing.T) {
	type args struct {
		ctx               context.Context
		g                 *blend4go.GalaxyInstance
		toolShedUrl       string
		name              string
		owner             string
		changesetRevision string
		removeFromDisk    bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Uninstall(tt.args.ctx, tt.args.g, tt.args.toolShedUrl, tt.args.name, tt.args.owner, tt.args.changesetRevision, tt.args.removeFromDisk); (err != nil) != tt.wantErr {
				t.Errorf("Uninstall() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUninstallID(t *testing.T) {
	type args struct {
		ctx            context.Context
		g              *blend4go.GalaxyInstance
		id             string
		removeFromDisk bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UninstallID(tt.args.ctx, tt.args.g, tt.args.id, tt.args.removeFromDisk); (err != nil) != tt.wantErr {
				t.Errorf("UninstallID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
