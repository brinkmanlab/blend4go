package tools

import "github.com/brinkmanlab/blend4go"

type ToolSection struct {
	Id         string  `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Version    string  `json:"version,omitempty"`
	ModelClass string  `json:"model_class,omitempty"`
	Elems      []*Tool `json:"elems,omitempty"`
}

func (t *ToolSection) GetID() blend4go.GalaxyID {
	return t.Id
}

func (t *ToolSection) SetID(id blend4go.GalaxyID) {
	t.Id = id
}

type Tool struct {
	galaxyInstance   *blend4go.GalaxyInstance
	Id               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Version          string   `json:"version,omitempty"`
	MinWidth         int      `json:"min_width,omitempty"`
	Target           string   `json:"target,omitempty"`
	Link             string   `json:"link,omitempty"`
	PanelSectionId   string   `json:"panel_section_id,omitempty"`
	EdamTopics       []string `json:"edam_topics,omitempty"`
	FormStyle        string   `json:"form_style,omitempty"`
	EdamOperations   []string `json:"edam_operations,omitempty"`
	Labels           []string `json:"labels,omitempty"`
	Description      string   `json:"description,omitempty"`
	ConfigFile       string   `json:"config_file,omitempty"`
	Xrefs            []string `json:"xrefs,omitempty"`
	PanelSectionName string   `json:"panel_section_name,omitempty"`
}

func (t *Tool) GetBasePath() string {
	return BasePath
}

func (t *Tool) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	t.galaxyInstance = g
}

func (t *Tool) GetID() blend4go.GalaxyID {
	return t.Id
}

func (t *Tool) SetID(id blend4go.GalaxyID) {
	t.Id = id
}

func (t *Tool) Reload() error {
	// GET /api/tools/{tool_id}/reload Reload specified tool
	panic("Implement me")
}

// Attempts to install requirements via the dependency resolver
func (t *Tool) InstallDependencies() error {
	// POST /api/tools/{tool_id}/dependencies
	panic("Implement me")
}

// DELETE /api/tools/{tool_id}/dependencies Attempts to uninstall requirements via the dependency resolver
func (t *Tool) UninstallDependencies() error {
	// POST /api/tools/{tool_id}/dependencies
	panic("Implement me")
}

//POST /api/tools/{tool_id}/build_dependency_cache Attempts to cache installed dependencies.
//GET /api/tools/{tool_id}/diagnostics Return diagnostic information to help debug panel and dependency related problems.

//GET /api/tools/{tool_id}/build Returns a tool model including dynamic parameters and updated values, repeats block etc.
//GET /api/tools/{tool_id}/test_data_path?tool_version={tool_version}
//GET /api/tools/{tool_id}/test_data_download?tool_version={tool_version}&filename={filename}
//GET /api/tools/{tool_id}/test_data?tool_version={tool_version}
//GET /api/tools/{tool_id}/requirements Return the resolver status for a specific tool id. [{“status”: “installed”, “name”: “hisat2”, “versionless”: false, “resolver_type”: “conda”, “version”: “2.0.3”, “type”: “package”}]
