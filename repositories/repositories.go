// repositories models represent and manipulate repositories within a Galaxy instance
// Relevant api endpoints are: `/api/tool_shed_repositories`
package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brinkmanlab/blend4go"
	"path"
	"strings"
	"time"
)

const BasePath = "/api/tool_shed_repositories"

var NON_TERMINAL_REPOSITORY_STATES = map[string]bool{
	"New":                                true,
	"Cloning":                            true,
	"Setting tool versions":              true,
	"Installing repository dependencies": true,
	"Installing tool dependencies":       true,
	"Loading proprietary datatypes":      true,
}

func List(ctx context.Context, g *blend4go.GalaxyInstance) ([]*Repository, error) {
	var repos []*Repository
	_, err := g.List(ctx, BasePath, &repos, nil)
	return repos, err
}

func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*Repository, error) {
	if res, err := g.Get(ctx, id, &Repository{}, nil); err == nil {
		return res.(*Repository), nil
	} else {
		return nil, err
	}
}

type repoInstallConfig struct {
	Id                            blend4go.GalaxyID `json:"id,omitempty"`
	ToolShedUrl                   string            `json:"tool_shed_url"`
	Name                          string            `json:"name"`
	Owner                         string            `json:"owner"`
	ChangesetRevision             string            `json:"changeset_revision"`
	InstallToolDependencies       bool              `json:"install_tool_dependencies,omitempty"`
	InstallRepositoryDependencies bool              `json:"install_repository_dependencies,omitempty"`
	InstallResolverDependencies   bool              `json:"install_resolver_dependencies,omitempty"`
	ToolPanelSectionId            blend4go.GalaxyID `json:"tool_panel_section_id,omitempty"`
	NewToolPanelSectionLabel      string            `json:"new_tool_panel_section_label,omitempty"`
	RemoveFromDisk                bool              `json:"remove_from_disk,omitempty"`
}

// Install a specified repository revision from a specified tool shed into Galaxy
func Install(ctx context.Context, g *blend4go.GalaxyInstance, toolShedUrl string, owner string, name string, changesetRevision string, installToolDependencies bool, installRepositoryDependencies bool, installResolverDependencies bool, toolPanelSectionId blend4go.GalaxyID, newToolPanelSectionLabel string, wait uint) ([]*Repository, error) {
	//https://github.com/galaxyproject/ephemeris/blob/474a1c1cd4d5444ece00a3e53eafcb234643db90/src/ephemeris/shed_tools.py#L374
	// POST /api/tool_shed_repositories/install_repository_revision
	// https://docs.galaxyproject.org/en/latest/api/api.html#galaxy.webapps.galaxy.api.tool_shed_repositories.ToolShedRepositoriesController.install_repository_revision
	if toolPanelSectionId != "" {
		// Ensure only one is non-empty
		newToolPanelSectionLabel = ""
	}
	// TODO changeset_revision == "" ? install latest
	if !strings.Contains(toolShedUrl, "://") {
		toolShedUrl = "https://" + toolShedUrl
	}
	config := repoInstallConfig{ToolShedUrl: toolShedUrl, Name: name, Owner: owner, ChangesetRevision: changesetRevision, InstallToolDependencies: installToolDependencies, InstallRepositoryDependencies: installRepositoryDependencies, InstallResolverDependencies: installResolverDependencies, ToolPanelSectionId: toolPanelSectionId, NewToolPanelSectionLabel: newToolPanelSectionLabel}
	if res, err := g.R(ctx).SetBody(config).Post(path.Join(BasePath, "new", "install_repository_revision")); err == nil {
		if res.StatusCode() == 504 && wait > 0 {
			// Endpoint timed out, try to recover
			if repos, err := List(ctx, g); err == nil {
				for _, repo := range repos {
					if repo.ToolShed == toolShedUrl && repo.Owner == owner && repo.Name == name && repo.ChangesetRevision == changesetRevision {
						_, pending := NON_TERMINAL_REPOSITORY_STATES[repo.Status]
						for pending {
							time.Sleep(time.Second)
							_, _ = g.Get(ctx, repo.Id, repo, &map[string]string{})
							_, pending = NON_TERMINAL_REPOSITORY_STATES[repo.Status]
							wait--
							if wait == 0 {
								break
							}
						}
						if pending {
							return nil, fmt.Errorf("%v took longer than %v seconds to install", path.Join(toolShedUrl, owner, name, changesetRevision), wait)
						} else {
							return []*Repository{repo}, nil
						}
					}
				}
				return nil, fmt.Errorf("could not recover from repository install timing out")
			} else {
				return nil, err
			}
		}
		if _, err := blend4go.HandleResponse(res); err == nil {
			if strings.HasPrefix(res.String(), "[") {
				repos := make([]*Repository, 0)
				if err := json.Unmarshal(res.Body(), &repos); err == nil {
					for _, repo := range repos {
						repo.SetGalaxyInstance(g)
					}
					return repos, nil
				} else {
					return nil, err
				}
			} else {
				status := &blend4go.StatusResponse{}
				if err := json.Unmarshal(res.Body(), status); err == nil {
					if status.Status == "ok" {
						return nil, nil
					} else {
						return nil, errors.New(status.Message)
					}
				} else {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// Uninstall a specified repository revision from a specified tool shed from Galaxy
func Uninstall(ctx context.Context, g *blend4go.GalaxyInstance, toolShedUrl string, owner string, name string, changesetRevision string, removeFromDisk bool) error {
	config := repoInstallConfig{ToolShedUrl: toolShedUrl, Name: name, Owner: owner, ChangesetRevision: changesetRevision, RemoveFromDisk: removeFromDisk}
	if res, err := g.R(ctx).SetBody(config).SetResult(blend4go.StatusResponse{}).Delete("/api/tool_shed_repositories/"); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			if result.(blend4go.StatusResponse).Status == "ok" {
				return errors.New(res.Result().(blend4go.StatusResponse).Message)
			}
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

// Uninstall a specified repository id
func UninstallID(ctx context.Context, g *blend4go.GalaxyInstance, id string, removeFromDisk bool) error {
	config := repoInstallConfig{RemoveFromDisk: removeFromDisk}
	if res, err := g.R(ctx).SetBody(config).SetResult(&blend4go.StatusResponse{}).Delete(path.Join(BasePath, id)); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			if result.(*blend4go.StatusResponse).Status == "ok" {
				return errors.New(res.Result().(*blend4go.StatusResponse).Message)
			}
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

// Check for updates to the specified repository, or all installed repositories.
func CheckForUpdates(ctx context.Context, g *blend4go.GalaxyInstance, repoID blend4go.GalaxyID) error {
	req := g.R(ctx)
	if repoID != "" {
		req.SetQueryParam("id", repoID)
	}
	if res, err := req.SetResult(blend4go.StatusResponse{}).Get("/api/tool_shed_repositories/check_for_updates"); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			if result.(blend4go.StatusResponse).Status != "ok" {
				return errors.New(res.Result().(blend4go.StatusResponse).Message)
			}
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func ResetMetadataAll(ctx context.Context, g *blend4go.GalaxyInstance) error {
	if res, err := g.R(ctx).SetResult(blend4go.StatusResponse{}).Put("/api/tool_shed_repositories/reset_metadata_on_installed_repositories"); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			if result.(blend4go.StatusResponse).Status != "ok" {
				return errors.New(res.Result().(blend4go.StatusResponse).Message)
			}
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}
