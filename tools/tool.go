package tools

import "blend4go"

type ToolSection struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Version    string `json:"version"`
	ModelClass string `json:"model_class"`
	Elems      []Tool `json:"elems"`
}

type Tool struct {
	galaxyInstance   *blend4go.GalaxyInstance
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	Version          string   `json:"version"`
	MinWidth         int      `json:"min_width"`
	Target           string   `json:"target"`
	Link             string   `json:"link"`
	PanelSectionId   string   `json:"panel_section_id"`
	EdamTopics       []string `json:"edam_topics"`
	FormStyle        string   `json:"form_style"`
	EdamOperations   []string `json:"edam_operations"`
	Labels           []string `json:"labels"`
	Description      string   `json:"description"`
	ConfigFile       string   `json:"config_file"`
	Xrefs            []string `json:"xrefs"`
	PanelSectionName string   `json:"panel_section_name"`
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
//POST /api/tools Execute tool with a given parameter payload
