package workflows

import (
	"github.com/brinkmanlab/blend4go"
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
	ToolId      string `json:"tool_id"`
	ToolVersion string `json:"tool_version"`
	Id          uint   `json:"id"`
	InputSteps  []StoredWorkflowInputStep
	// ToolInputs	? `json:"tool_inputs"` TODO?
	Type       string `json:"type"`
	Annotation string `json:"annotation"`
}

type StoredWorkflow struct {
	galaxyInstance     *blend4go.GalaxyInstance
	Id                 blend4go.GalaxyID        `json:"id"`
	Name               string                   `json:"name"`
	Tags               []string                 `json:"tags"`
	Deleted            bool                     `json:"deleted"`
	LatestWorkflowUuid string                   `json:"latest_workflow_uuid"`
	ShowInToolPanel    bool                     `json:"show_in_tool_panel"`
	Url                string                   `json:"url"`
	NumberOfSteps      uint                     `json:"number_of_steps"`
	Published          bool                     `json:"published"`
	Owner              string                   `json:"owner"`
	ModelClass         string                   `json:"model_class"`
	Inputs             []StoredWorkflowInput    `json:"inputs"`
	Annotation         string                   `json:"annotation"`
	Version            uint                     `json:"version"`
	Steps              []WorkflowInvocationStep `json:"steps"`
}

func (w *StoredWorkflow) GetBasePath() string {
	return BasePath
}

func (w *StoredWorkflow) SetGalaxyInstance(g *blend4go.GalaxyInstance) {
	w.galaxyInstance = g
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
