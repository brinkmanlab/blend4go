package workflows

import (
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/repositories"
	"path"
)

type StoredWorkflowInput struct {
	Uuid  string `json:"uuid"`
	Value string `json:"value"`
	Label string `json:"label"`
}

type StoredWorkflowInputStep struct {
	StepOutput string `json:"step_output"`
	SourceStep uint   `json:"source_step"`
}

type StoredWorkflowStep struct {
	ToolId      string                    `json:"tool_id,omitempty"`
	ToolVersion string                    `json:"tool_version,omitempty"`
	Id          uint                      `json:"id,omitempty"`
	InputSteps  []StoredWorkflowInputStep `json:"input_steps,omitempty"`
	// ToolInputs	? `json:"tool_inputs,omitempty"` TODO?
	Type       string `json:"type,omitempty"`
	Annotation string `json:"annotation,omitempty"`
}

type StoredWorkflow struct {
	galaxyInstance     *blend4go.GalaxyInstance
	Id                 blend4go.GalaxyID        `json:"id,omitempty"`
	Name               string                   `json:"name,omitempty"`
	Tags               []string                 `json:"tags,omitempty"`
	Deleted            bool                     `json:"deleted,omitempty"`
	LatestWorkflowUuid string                   `json:"latest_workflow_uuid,omitempty"`
	ShowInToolPanel    bool                     `json:"show_in_tool_panel,omitempty"`
	Url                string                   `json:"url,omitempty"`
	NumberOfSteps      uint                     `json:"number_of_steps,omitempty"`
	Published          bool                     `json:"published,omitempty"`
	Owner              string                   `json:"owner,omitempty"`
	ModelClass         string                   `json:"model_class,omitempty"`
	Inputs             []StoredWorkflowInput    `json:"inputs,omitempty"`
	Annotation         string                   `json:"annotation,omitempty"`
	Version            uint                     `json:"version,omitempty"`
	Steps              []WorkflowInvocationStep `json:"steps,omitempty"`
}

func (w *StoredWorkflow) GetBasePath() string {
	return BasePath
}

func (w *StoredWorkflow) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	w.galaxyInstance = g
}

func (w *StoredWorkflow) GetID() blend4go.GalaxyID {
	return w.Id
}

func (w *StoredWorkflow) SetID(id blend4go.GalaxyID) {
	w.Id = id
}

func NewStoredWorkflow(g *blend4go.GalaxyInstance, json string) (*StoredWorkflow, error) {
	if res, err := g.R().SetResult(&StoredWorkflow{galaxyInstance: g}).SetBody(map[string]string{
		"workflow": json,
	}).Post(BasePath); err == nil {
		return res.Result().(*StoredWorkflow), nil
	} else {
		return nil, err
	}
}

// GET /api/workflows/{encoded_workflow_id}/versions
// instance (boolean) – true if fetch by Workflow ID instead of StoredWorkflow id, false by default.

// GET /api/workflows/{encoded_workflow_id}/download

// Deletes a specified workflow
func (w *StoredWorkflow) Delete() error {
	// DELETE /api/workflows/{encoded_workflow_id}
	return w.galaxyInstance.Delete(w.Id, w)
}

// PUT /api/workflows/{id}
func (w *StoredWorkflow) Update() error {
	return w.galaxyInstance.Put(w.Id, w)
}

// Schedule the workflow specified by workflow_id to run.
func (w *StoredWorkflow) Invoke() error {
	// POST /api/workflows/{encoded_workflow_id}/invocations
	panic("Implement me") // TODO
}
