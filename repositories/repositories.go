package repositories

import (
	"context"
	"errors"
	"github.com/brinkmanlab/blend4go"
)

const BasePath = "/api/tool_shed_repositories"

func List(ctx context.Context, g *blend4go.GalaxyInstance) ([]Repository, error) {
	if res, err := g.List(ctx, BasePath, []Repository{}, nil); err == nil {
		return res.([]Repository), nil
	} else {
		return nil, err
	}
}

func Get(ctx context.Context, g *blend4go.GalaxyInstance, id blend4go.GalaxyID) (*Repository, error) {
	if res, err := g.Get(ctx, id, &Repository{}); err == nil {
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
func Install(ctx context.Context, g *blend4go.GalaxyInstance, toolShedUrl string, name string, owner string, changesetRevision string, installToolDependencies bool, installRepositoryDependencies bool, installResolverDependencies bool, toolPanelSectionId blend4go.GalaxyID, newToolPanelSectionLabel string) error {
	//https://github.com/galaxyproject/ephemeris/blob/474a1c1cd4d5444ece00a3e53eafcb234643db90/src/ephemeris/shed_tools.py#L374
	// POST /api/tool_shed_repositories/install_repository_revision
	// https://docs.galaxyproject.org/en/latest/api/api.html#galaxy.webapps.galaxy.api.tool_shed_repositories.ToolShedRepositoriesController.install_repository_revision
	if toolPanelSectionId != "" {
		// Ensure only one is non-empty
		newToolPanelSectionLabel = ""
	}
	config := repoInstallConfig{ToolShedUrl: toolShedUrl, Name: name, Owner: owner, ChangesetRevision: changesetRevision, InstallToolDependencies: installToolDependencies, InstallRepositoryDependencies: installRepositoryDependencies, InstallResolverDependencies: installResolverDependencies, ToolPanelSectionId: toolPanelSectionId, NewToolPanelSectionLabel: newToolPanelSectionLabel}
	if res, err := g.R(ctx).SetBody(config).SetResult(blend4go.StatusResponse{}).Post("/api/tool_shed_repositories/install_repository_revision"); err == nil {
		if res.Result().(blend4go.StatusResponse).Status == "ok" {
			return errors.New(res.Result().(blend4go.StatusResponse).Message)
		}
	} else {
		return err
	}
	return nil
}

// Uninstall a specified repository revision from a specified tool shed from Galaxy
func Uninstall(ctx context.Context, g *blend4go.GalaxyInstance, toolShedUrl string, name string, owner string, changesetRevision string, removeFromDisk bool) error {
	config := repoInstallConfig{ToolShedUrl: toolShedUrl, Name: name, Owner: owner, ChangesetRevision: changesetRevision, RemoveFromDisk: removeFromDisk}
	if res, err := g.R(ctx).SetBody(config).SetResult(blend4go.StatusResponse{}).Delete("/api/tool_shed_repositories/"); err == nil {
		if res.Result().(blend4go.StatusResponse).Status == "ok" {
			return errors.New(res.Result().(blend4go.StatusResponse).Message)
		}
	} else {
		return err
	}
	return nil
}

// Uninstall a specified repository id
func UninstallID(ctx context.Context, g *blend4go.GalaxyInstance, id string, removeFromDisk bool) error {
	config := repoInstallConfig{Id: id, RemoveFromDisk: removeFromDisk}
	if res, err := g.R(ctx).SetBody(config).SetResult(blend4go.StatusResponse{}).Delete("/api/tool_shed_repositories/"); err == nil {
		if res.Result().(blend4go.StatusResponse).Status == "ok" {
			return errors.New(res.Result().(blend4go.StatusResponse).Message)
		}
	} else {
		return err
	}
	return nil
}

// Check for updates to the specified repository, or all installed repositories.
func CheckForUpdates(ctx context.Context, g *blend4go.GalaxyInstance, repoID blend4go.GalaxyID) error {
	req := g.R(ctx)
	if repoID != "" {
		req.SetQueryParam("id", repoID)
	}
	if res, err := req.SetResult(blend4go.StatusResponse{}).Get("/api/tool_shed_repositories/check_for_updates"); err == nil {
		if res.Result().(blend4go.StatusResponse).Status != "ok" {
			return errors.New(res.Result().(blend4go.StatusResponse).Message)
		}
		return nil
	} else {
		return err
	}
}

func ResetMetadataAll(ctx context.Context, g *blend4go.GalaxyInstance) error {
	if res, err := g.R(ctx).SetResult(blend4go.StatusResponse{}).Get("/api/tool_shed_repositories/check_for_updates"); err == nil {
		if res.Result().(blend4go.StatusResponse).Status != "ok" {
			return errors.New(res.Result().(blend4go.StatusResponse).Message)
		}
		return nil
	} else {
		return err
	}
}
