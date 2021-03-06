package workflows

import (
	"context"
	"encoding/json"
	"github.com/brinkmanlab/blend4go"
	"github.com/brinkmanlab/blend4go/repositories"
	"path"
)

type StoredWorkflowInput struct {
	Uuid  string `json:"uuid,omitempty"`
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}

type StoredWorkflowInputStep struct {
	StepOutput string `json:"step_output,omitempty"`
	SourceStep uint   `json:"source_step,omitempty"`
}

type StoredWorkflowStep struct {
	ToolId      string                              `json:"tool_id,omitempty"`
	ToolVersion string                              `json:"tool_version,omitempty"`
	Id          uint                                `json:"id,omitempty"`
	InputSteps  map[string]*StoredWorkflowInputStep `json:"input_steps,omitempty"`
	// ToolInputs	? `json:"tool_inputs,omitempty"` TODO?
	Type       string `json:"type,omitempty"`
	Annotation string `json:"annotation,omitempty"`
}

type StoredWorkflow struct {
	galaxyInstance     *blend4go.GalaxyInstance
	Id                 blend4go.GalaxyID               `json:"id,omitempty"`
	Name               string                          `json:"name,omitempty"`
	Tags               []string                        `json:"tags,omitempty"`
	Deleted            bool                            `json:"deleted,omitempty"`
	LatestWorkflowUuid string                          `json:"latest_workflow_uuid,omitempty"`
	ShowInToolPanel    bool                            `json:"show_in_tool_panel,omitempty"`
	Url                string                          `json:"url,omitempty"`
	NumberOfSteps      uint                            `json:"number_of_steps,omitempty"`
	Published          bool                            `json:"published,omitempty"`
	Owner              string                          `json:"owner,omitempty"`
	ModelClass         string                          `json:"model_class,omitempty"`
	Inputs             map[string]*StoredWorkflowInput `json:"inputs,omitempty"`
	Annotation         string                          `json:"annotation,omitempty"`
	Version            uint                            `json:"version,omitempty"`
	Steps              map[string]*StoredWorkflowStep  `json:"steps,omitempty"`
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

func NewStoredWorkflow(ctx context.Context, g *blend4go.GalaxyInstance, j string, importTools, publish, importable bool) (*StoredWorkflow, error) {
	jd := &map[string]interface{}{}
	err := json.Unmarshal([]byte(j), jd)
	if err != nil {
		return nil, err
	}
	body := map[string]interface{}{
		"workflow": jd,
	}
	if importTools {
		body["import_tools"] = "true"
	}
	if publish {
		body["publish"] = "true"
	}
	if importable {
		body["importable"] = "true"
	}
	if res, err := g.R(ctx).SetResult(&StoredWorkflow{galaxyInstance: g}).SetBody(body).Post(BasePath); err == nil {
		if result, err := blend4go.HandleResponse(res); err == nil {
			return result.(*StoredWorkflow), nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

// GET /api/workflows/{encoded_workflow_id}/versions
// instance (boolean) – true if fetch by Workflow ID instead of StoredWorkflow id, false by default.

// Delete a specified workflow
func (w *StoredWorkflow) Delete(ctx context.Context) error {
	// DELETE /api/workflows/{encoded_workflow_id}
	return w.galaxyInstance.Delete(ctx, w, nil)
}

// Update the specified workflow. If json == "", only the name, annotation, and show_in_tool_panel will be updated.
func (w *StoredWorkflow) Update(ctx context.Context, j string) error {
	// TODO https://github.com/galaxyproject/galaxy/issues/10682
	// TODO https://github.com/galaxyproject/galaxy/issues/10683
	// TODO https://github.com/galaxyproject/galaxy/issues/10684
	// PUT /api/workflows/{id}
	body := make(map[string]interface{})
	if w.ShowInToolPanel {
		body["menu_entry"] = "True"
	} else {
		body["menu_entry"] = "False"
	}

	if j != "" {
		jd := &map[string]interface{}{}
		err := json.Unmarshal([]byte(j), jd)
		if err != nil {
			return err
		}
		body["workflow"] = jd
	}

	body["name"] = w.Name
	body["annotation"] = w.Annotation

	if res, err := w.galaxyInstance.R(ctx).SetResult(w).SetBody(body).Put(path.Join(w.GetBasePath(), w.GetID())); err == nil {
		if _, err := blend4go.HandleResponse(res); err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (w *StoredWorkflow) Download(ctx context.Context) (string, error) {
	// GET /api/workflows/{encoded_workflow_id}/download
	if res, err := w.galaxyInstance.R(ctx).Get(path.Join(w.GetBasePath(), w.GetID(), "download")); err == nil {
		if _, err := blend4go.HandleResponse(res); err == nil {
			return string(res.Body()), nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}

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
