package workflows

import (
	"context"
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
	ToolId      string                     `json:"tool_id,omitempty"`
	ToolVersion string                     `json:"tool_version,omitempty"`
	Id          uint                       `json:"id,omitempty"`
	InputSteps  []*StoredWorkflowInputStep `json:"input_steps,omitempty"`
	// ToolInputs	? `json:"tool_inputs,omitempty"` TODO?
	Type       string `json:"type,omitempty"`
	Annotation string `json:"annotation,omitempty"`
}

type StoredWorkflow struct {
	galaxyInstance     *blend4go.GalaxyInstance
	Id                 blend4go.GalaxyID         `json:"id,omitempty"`
	Name               string                    `json:"name,omitempty"`
	Tags               []string                  `json:"tags,omitempty"`
	Deleted            bool                      `json:"deleted,omitempty"`
	LatestWorkflowUuid string                    `json:"latest_workflow_uuid,omitempty"`
	ShowInToolPanel    bool                      `json:"show_in_tool_panel,omitempty"`
	Url                string                    `json:"url,omitempty"`
	NumberOfSteps      uint                      `json:"number_of_steps,omitempty"`
	Published          bool                      `json:"published,omitempty"`
	Owner              string                    `json:"owner,omitempty"`
	ModelClass         string                    `json:"model_class,omitempty"`
	Inputs             []*StoredWorkflowInput    `json:"inputs,omitempty"`
	Annotation         string                    `json:"annotation,omitempty"`
	Version            uint                      `json:"version,omitempty"`
	Steps              []*WorkflowInvocationStep `json:"steps,omitempty"`
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

func NewStoredWorkflow(ctx context.Context, g *blend4go.GalaxyInstance, json string) (*StoredWorkflow, error) {
	if res, err := g.R(ctx).SetResult(&StoredWorkflow{galaxyInstance: g}).SetBody(map[string]string{
		"workflow": json,
	}).Post(BasePath); err == nil {
		return res.Result().(*StoredWorkflow), nil
	} else {
		return nil, err
	}
}

// GET /api/workflows/{encoded_workflow_id}/versions
// instance (boolean) – true if fetch by Workflow ID instead of StoredWorkflow id, false by default.

// Delete a specified workflow
func (w *StoredWorkflow) Delete(ctx context.Context) error {
	// DELETE /api/workflows/{encoded_workflow_id}
	return w.galaxyInstance.Delete(ctx, w)
}

// Update the specified workflow. If json == "", only the name, annotation, and show_in_tool_panel will be updated.
func (w *StoredWorkflow) Update(ctx context.Context, json string) error {
	// PUT /api/workflows/{id}
	body := make(map[string]string)
	if w.ShowInToolPanel {
		body["menu_entry"] = "True"
	} else {
		body["menu_entry"] = "False"
	}

	if json != "" {
		body["workflow"] = json
	}

	body["name"] = w.Name
	body["annotation"] = w.Annotation

	_, err := w.galaxyInstance.R(ctx).SetResult(w).SetBody(body).Put(path.Join(w.GetBasePath(), w.GetID()))
	return err
}

func (w *StoredWorkflow) Download(ctx context.Context) (string, error) {
	// GET /api/workflows/{encoded_workflow_id}/download
	res, err := w.galaxyInstance.R(ctx).Get(path.Join(w.GetBasePath(), w.GetID(), "download"))
	return res.String(), err
}

func (w *StoredWorkflow) Repositories(ctx context.Context) ([]*repositories.Repository, error) {
	if workflow, err := w.Download(ctx); err == nil {
		return Repositories(workflow)
	} else {
		return nil, err
	}
}

// Schedule the workflow specified by workflow_id to run.
func (w *StoredWorkflow) Invoke(ctx context.Context) error {
	// POST /api/workflows/{encoded_workflow_id}/invocations
	panic("Implement me") // TODO
}
