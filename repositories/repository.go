package repositories

import (
	"context"
	"fmt"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/tools"
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

// list Tools provided by Repository
func (r *Repository) Tools(ctx context.Context) ([]*tools.Tool, error) {
	if res, err := r.galaxyInstance.R(ctx).SetQueryParams(map[string]string{
		"tool_shed_url": "https://" + r.ToolShed + "/",
		"id":            r.GetID(),
		"controller":    "repositories",
		"action":        "metadata",
	}).SetResult(map[string]interface{}{}).Get("/api/tool_shed/request"); err == nil {
		if result, err := r.galaxyInstance.HandleResponse(res); err == nil {
			changesets := *result.(*map[string]interface{})
			if meta, ok := changesets[fmt.Sprintf("%v:%v", r.CtxRev, r.ChangesetRevision)]; ok {
				if t, ok := (meta.(map[string]interface{}))["tools"]; ok {
					toolList := t.([]interface{})
					toolModels := make([]*tools.Tool, len(toolList), len(toolList))
					for i, item := range toolList {
						tool := item.(map[string]interface{})
						toolModels[i] = &tools.Tool{
							Id:          tool["id"].(string),
							Guid:        tool["guid"].(string),
							Name:        tool["name"].(string),
							Version:     tool["version"].(string),
							Description: tool["description"].(string),
							ConfigFile:  tool["tool_config"].(string),
						}
					}
					return toolModels, nil
				} else {
					return nil, fmt.Errorf("unexpected response body returned from API: %v", meta)
				}
			} else {
				return nil, fmt.Errorf("unexpected changeset returned from API: %v", changesets)
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
