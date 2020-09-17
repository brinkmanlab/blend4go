package repositories

import (
	"github.com/brinkmanlab/blend4go"
)

type Repository struct {
	galaxyInstance             *blend4go.GalaxyInstance
	Id                         blend4go.GalaxyID `json:"id,omitempty"`
	Status                     string            `json:"status,omitempty"`
	Name                       string            `json:"name"`
	Deleted                    bool              `json:"deleted,omitempty"`
	CtxRev                     string            `json:"ctx_rev,omitempty"`
	ErrorMessage               string            `json:"error_message,omitempty"`
	InstalledChangesetRevision string            `json:"installed_changeset_revision,omitempty"`
	ToolShed                   string            `json:"tool_shed"`
	DistToShed                 bool              `json:"dist_to_shed,omitempty"`
	Url                        string            `json:"url,omitempty"`
	Uninstalled                bool              `json:"uninstalled,omitempty"`
	Owner                      string            `json:"owner"`
	ChangesetRevision          string            `json:"changeset_revision"`
	IncludeDatatypes           bool              `json:"include_datatypes,omitempty"`
	ToolShedStatus             struct {
		LatestInstallableRevision string `json:"latest_installable_revision"`
		RevisionUpdate            string `json:"revision_update"`
		RevisionUpgrade           string `json:"revision_upgrade"`
		RepositoryDeprecated      string `json:"repository_deprecated"`
	} `json:"tool_shed_status,omitempty"`
}

func (r *Repository) GetBasePath() string {
	return BasePath
}

func (r *Repository) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	r.galaxyInstance = g
}

func (r *Repository) GetID() blend4go.GalaxyID {
	return r.Id
}

func (r *Repository) SetID(id blend4go.GalaxyID) {
	r.Id = id
}

func (r *Repository) Reload() error {
	panic("Implement me")
}
