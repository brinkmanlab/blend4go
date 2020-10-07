package repositories

import (
	"github.com/brinkmanlab/blend4go"
	"testing"
)

func TestRepository_GetBasePath(t *testing.T) {
	type fields struct {
		galaxyInstance             *blend4go.GalaxyInstance
		Id                         blend4go.GalaxyID
		Status                     string
		Name                       string
		Deleted                    bool
		CtxRev                     string
		ErrorMessage               string
		InstalledChangesetRevision string
		ToolShed                   string
		DistToShed                 bool
		Url                        string
		Uninstalled                bool
		Owner                      string
		ChangesetRevision          string
		IncludeDatatypes           bool
		ToolShedStatus             struct {
			LatestInstallableRevision string `json:"latest_installable_revision"`
			RevisionUpdate            string `json:"revision_update"`
			RevisionUpgrade           string `json:"revision_upgrade"`
			RepositoryDeprecated      string `json:"repository_deprecated"`
		}
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
			r := &Repository{
				galaxyInstance:             tt.fields.galaxyInstance,
				Id:                         tt.fields.Id,
				Status:                     tt.fields.Status,
				Name:                       tt.fields.Name,
				Deleted:                    tt.fields.Deleted,
				CtxRev:                     tt.fields.CtxRev,
				ErrorMessage:               tt.fields.ErrorMessage,
				InstalledChangesetRevision: tt.fields.InstalledChangesetRevision,
				ToolShed:                   tt.fields.ToolShed,
				DistToShed:                 tt.fields.DistToShed,
				Url:                        tt.fields.Url,
				Uninstalled:                tt.fields.Uninstalled,
				Owner:                      tt.fields.Owner,
				ChangesetRevision:          tt.fields.ChangesetRevision,
				IncludeDatatypes:           tt.fields.IncludeDatatypes,
				ToolShedStatus:             tt.fields.ToolShedStatus,
			}
			if got := r.GetBasePath(); got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetID(t *testing.T) {
	type fields struct {
		galaxyInstance             *blend4go.GalaxyInstance
		Id                         blend4go.GalaxyID
		Status                     string
		Name                       string
		Deleted                    bool
		CtxRev                     string
		ErrorMessage               string
		InstalledChangesetRevision string
		ToolShed                   string
		DistToShed                 bool
		Url                        string
		Uninstalled                bool
		Owner                      string
		ChangesetRevision          string
		IncludeDatatypes           bool
		ToolShedStatus             struct {
			LatestInstallableRevision string `json:"latest_installable_revision"`
			RevisionUpdate            string `json:"revision_update"`
			RevisionUpgrade           string `json:"revision_upgrade"`
			RepositoryDeprecated      string `json:"repository_deprecated"`
		}
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
			r := &Repository{
				galaxyInstance:             tt.fields.galaxyInstance,
				Id:                         tt.fields.Id,
				Status:                     tt.fields.Status,
				Name:                       tt.fields.Name,
				Deleted:                    tt.fields.Deleted,
				CtxRev:                     tt.fields.CtxRev,
				ErrorMessage:               tt.fields.ErrorMessage,
				InstalledChangesetRevision: tt.fields.InstalledChangesetRevision,
				ToolShed:                   tt.fields.ToolShed,
				DistToShed:                 tt.fields.DistToShed,
				Url:                        tt.fields.Url,
				Uninstalled:                tt.fields.Uninstalled,
				Owner:                      tt.fields.Owner,
				ChangesetRevision:          tt.fields.ChangesetRevision,
				IncludeDatatypes:           tt.fields.IncludeDatatypes,
				ToolShedStatus:             tt.fields.ToolShedStatus,
			}
			if got := r.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Reload(t *testing.T) {
	type fields struct {
		galaxyInstance             *blend4go.GalaxyInstance
		Id                         blend4go.GalaxyID
		Status                     string
		Name                       string
		Deleted                    bool
		CtxRev                     string
		ErrorMessage               string
		InstalledChangesetRevision string
		ToolShed                   string
		DistToShed                 bool
		Url                        string
		Uninstalled                bool
		Owner                      string
		ChangesetRevision          string
		IncludeDatatypes           bool
		ToolShedStatus             struct {
			LatestInstallableRevision string `json:"latest_installable_revision"`
			RevisionUpdate            string `json:"revision_update"`
			RevisionUpgrade           string `json:"revision_upgrade"`
			RepositoryDeprecated      string `json:"repository_deprecated"`
		}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				galaxyInstance:             tt.fields.galaxyInstance,
				Id:                         tt.fields.Id,
				Status:                     tt.fields.Status,
				Name:                       tt.fields.Name,
				Deleted:                    tt.fields.Deleted,
				CtxRev:                     tt.fields.CtxRev,
				ErrorMessage:               tt.fields.ErrorMessage,
				InstalledChangesetRevision: tt.fields.InstalledChangesetRevision,
				ToolShed:                   tt.fields.ToolShed,
				DistToShed:                 tt.fields.DistToShed,
				Url:                        tt.fields.Url,
				Uninstalled:                tt.fields.Uninstalled,
				Owner:                      tt.fields.Owner,
				ChangesetRevision:          tt.fields.ChangesetRevision,
				IncludeDatatypes:           tt.fields.IncludeDatatypes,
				ToolShedStatus:             tt.fields.ToolShedStatus,
			}
			if err := r.Reload(); (err != nil) != tt.wantErr {
				t.Errorf("Reload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_SetGalaxyInstance(t *testing.T) {
	type fields struct {
		galaxyInstance             *blend4go.GalaxyInstance
		Id                         blend4go.GalaxyID
		Status                     string
		Name                       string
		Deleted                    bool
		CtxRev                     string
		ErrorMessage               string
		InstalledChangesetRevision string
		ToolShed                   string
		DistToShed                 bool
		Url                        string
		Uninstalled                bool
		Owner                      string
		ChangesetRevision          string
		IncludeDatatypes           bool
		ToolShedStatus             struct {
			LatestInstallableRevision string `json:"latest_installable_revision"`
			RevisionUpdate            string `json:"revision_update"`
			RevisionUpgrade           string `json:"revision_upgrade"`
			RepositoryDeprecated      string `json:"repository_deprecated"`
		}
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
			r := &Repository{
				galaxyInstance:             tt.fields.galaxyInstance,
				Id:                         tt.fields.Id,
				Status:                     tt.fields.Status,
				Name:                       tt.fields.Name,
				Deleted:                    tt.fields.Deleted,
				CtxRev:                     tt.fields.CtxRev,
				ErrorMessage:               tt.fields.ErrorMessage,
				InstalledChangesetRevision: tt.fields.InstalledChangesetRevision,
				ToolShed:                   tt.fields.ToolShed,
				DistToShed:                 tt.fields.DistToShed,
				Url:                        tt.fields.Url,
				Uninstalled:                tt.fields.Uninstalled,
				Owner:                      tt.fields.Owner,
				ChangesetRevision:          tt.fields.ChangesetRevision,
				IncludeDatatypes:           tt.fields.IncludeDatatypes,
				ToolShedStatus:             tt.fields.ToolShedStatus,
			}
		})
	}
}

func TestRepository_SetID(t *testing.T) {
	type fields struct {
		galaxyInstance             *blend4go.GalaxyInstance
		Id                         blend4go.GalaxyID
		Status                     string
		Name                       string
		Deleted                    bool
		CtxRev                     string
		ErrorMessage               string
		InstalledChangesetRevision string
		ToolShed                   string
		DistToShed                 bool
		Url                        string
		Uninstalled                bool
		Owner                      string
		ChangesetRevision          string
		IncludeDatatypes           bool
		ToolShedStatus             struct {
			LatestInstallableRevision string `json:"latest_installable_revision"`
			RevisionUpdate            string `json:"revision_update"`
			RevisionUpgrade           string `json:"revision_upgrade"`
			RepositoryDeprecated      string `json:"repository_deprecated"`
		}
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
			r := &Repository{
				galaxyInstance:             tt.fields.galaxyInstance,
				Id:                         tt.fields.Id,
				Status:                     tt.fields.Status,
				Name:                       tt.fields.Name,
				Deleted:                    tt.fields.Deleted,
				CtxRev:                     tt.fields.CtxRev,
				ErrorMessage:               tt.fields.ErrorMessage,
				InstalledChangesetRevision: tt.fields.InstalledChangesetRevision,
				ToolShed:                   tt.fields.ToolShed,
				DistToShed:                 tt.fields.DistToShed,
				Url:                        tt.fields.Url,
				Uninstalled:                tt.fields.Uninstalled,
				Owner:                      tt.fields.Owner,
				ChangesetRevision:          tt.fields.ChangesetRevision,
				IncludeDatatypes:           tt.fields.IncludeDatatypes,
				ToolShedStatus:             tt.fields.ToolShedStatus,
			}
		})
	}
}
